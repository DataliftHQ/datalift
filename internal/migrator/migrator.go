package migrator

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"

	gatewayv1 "go.datalift.io/datalift/internal/config/gateway/v1"
	postgresv1 "go.datalift.io/datalift/internal/config/service/db/postgres/v1"
	"go.datalift.io/datalift/internal/gateway"
	pgservice "go.datalift.io/datalift/internal/service/db/postgres"
)

type Migrator struct {
	log          *zap.Logger
	force        bool
	sourceDriver source.Driver
	config       *gatewayv1.Config
}

func New(log *zap.Logger, configPath string, sourceInstance source.Driver, force bool) *Migrator {
	cfg := gateway.MustReadOrValidateConfig(
		&gateway.Flags{
			ConfigPath: configPath,
		},
	)

	return &Migrator{
		log:          log,
		config:       cfg,
		force:        force,
		sourceDriver: sourceInstance,
	}
}

func (m *Migrator) Up() {
	sqlMigrate := m.setupSqlMigrator()

	msg := "migration has the potential to cause irrevocable data loss, verify information above"
	m.confirmWithUser(msg)

	m.log.Info("applying up migrations")
	err := sqlMigrate.Up()
	if err != nil && err != migrate.ErrNoChange {
		m.log.Fatal("failed running migrations", zap.Error(err))
	}
}

func (m *Migrator) Down() {
	sqlMigrate := m.setupSqlMigrator()
	version, _, err := sqlMigrate.Version()
	if err != nil {
		m.log.Fatal("failed to acquire migration version", zap.Error(err))
	}

	msg := fmt.Sprintf(
		"Migrating DOWN by ONE version from (%d -> %d) this migration has the potential to cause irrevocable data loss, verify host information above",
		version, (version - 1))

	m.confirmWithUser(msg)

	// Migrate back by 1 change
	m.log.Info("applying migrations down")
	err = sqlMigrate.Steps(-1)
	if err != nil && err != migrate.ErrNoChange {
		m.log.Fatal("failed running migrations", zap.Error(err))
	}
}

func (m *Migrator) setupSqlMigrator() *migrate.Migrate {
	sqlDB, _ := m.setupSqlClient()

	// Ping database and bring up driver.
	if err := sqlDB.Ping(); err != nil {
		m.log.Fatal("error pinging db", zap.Error(err))
	}

	// Determine version storage table.
	cfg := &postgres.Config{
		MigrationsTable: postgres.DefaultMigrationsTable,
	}
	m.log.Info("using migration table", zap.String("migrationTable", cfg.MigrationsTable))

	// Create driver.
	driver, err := postgres.WithInstance(sqlDB, cfg)
	if err != nil {
		m.log.Fatal("error creating pg driver", zap.Error(err))
	}

	sqlMigrate, err := migrate.NewWithInstance(
		"iofs",
		m.sourceDriver,
		"postgres",
		driver,
	)
	if err != nil {
		m.log.Fatal("error creating migrator", zap.Error(err))
	}

	return sqlMigrate
}

func (m *Migrator) setupSqlClient() (*sql.DB, string) {
	var sqlDB *sql.DB
	var hostInfo string

	for _, s := range m.config.Services {
		if s.Name == pgservice.Name {
			pgdb, err := pgservice.New(s.TypedConfig, m.log, tally.NoopScope)
			if err != nil {
				m.log.Fatal("error creating db", zap.Error(err))
			}

			cfg := &postgresv1.Config{}
			if err := s.TypedConfig.UnmarshalTo(cfg); err != nil {
				m.log.Fatal("could not convert config", zap.Error(err))
			}

			sqlDB = pgdb.(pgservice.Client).DB()
			hostInfo = fmt.Sprintf("%s@%s:%d", cfg.Connection.User, cfg.Connection.Host, cfg.Connection.Port)

			break
		}
	}
	if sqlDB == nil {
		m.log.Fatal("no database found in config")
	}

	return sqlDB, hostInfo
}

func (m *Migrator) confirmWithUser(msg string) {
	_, hostInfo := m.setupSqlClient()
	// Verify that user wants to continue (unless -f for force is passed as a flag).
	m.log.Info("using database", zap.String("hostInfo", hostInfo))
	if !m.force {
		m.log.Warn(msg)

		fmt.Printf("\n*** Continue with migration? [y/N] ")
		var answer string
		if _, err := fmt.Scanln(&answer); err != nil && err.Error() != "unexpected newline" {
			m.log.Fatal("could not read user input", zap.Error(err))
		}
		if strings.ToLower(answer) != "y" {
			m.log.Fatal("aborting, enter 'y' to continue or use the '-f' (force) option")
		}
		fmt.Println()
	}
}
