package migrate

import (
	"embed"
	_ "embed"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/spf13/cobra"
	"go.datalift.io/datalift/internal/migrator"
	"go.uber.org/zap"
)

//go:embed migrations/*.sql
var fs embed.FS

type Cmd struct {
	Cmd *cobra.Command
}

func NewCmd(log *zap.Logger) *Cmd {
	root := &Cmd{}
	cmd := &cobra.Command{
		Use:           "migrate",
		Short:         "Migrates database schema",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.NoArgs,
	}

	cmd.AddCommand(
		newUpCmd(log),
		newDownCmd(log),
	)
	root.Cmd = cmd
	return root
}

func newUpCmd(log *zap.Logger) *cobra.Command {
	var config string
	var force bool

	cmd := &cobra.Command{
		Use:           "up",
		Short:         "Apply database schema",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.NoArgs,
		Run: func(_ *cobra.Command, _ []string) {
			sd, err := iofs.New(fs, "migrations")
			if err != nil {
				log.Fatal("error creating source driver from embedded assets", zap.Error(err))
			}

			m := migrator.New(log, config, sd, force)
			m.Up()
		},
	}

	cmd.Flags().StringVarP(&config, "config", "c", "datalift.yaml", "Path to YAML configuration")
	cmd.Flags().BoolVarP(&force, "force", "f", false, "Attempt to migrate scheme without prompting for confirmation")

	return cmd
}

func newDownCmd(log *zap.Logger) *cobra.Command {
	var config string
	var force bool

	cmd := &cobra.Command{
		Use:           "down",
		Short:         "Downgrade database schema by one version",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.NoArgs,
		Run: func(_ *cobra.Command, _ []string) {
			sd, err := iofs.New(fs, "migrations")
			if err != nil {
				log.Fatal("error creating source driver from embedded assets", zap.Error(err))
			}

			m := migrator.New(log, config, sd, force)
			m.Down()
		},
	}

	cmd.Flags().StringVarP(&config, "config", "c", "datalift.yaml", "Path to YAML configuration")
	cmd.Flags().BoolVarP(&force, "force", "f", false, "Attempt to migrate scheme without prompting for confirmation")

	return cmd
}
