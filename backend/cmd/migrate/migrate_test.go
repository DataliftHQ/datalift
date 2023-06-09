package main

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"

	"github.com/DataliftHQ/datalift/backend/gateway"
)

func TestSetupSqlClient(t *testing.T) {
	log := zaptest.NewLogger(t)

	pwd, err := os.Getwd()
	assert.NoError(t, err)

	flags := &gateway.Flags{
		ConfigPath: path.Join(pwd, "testdata/datalift-config-test.yaml"),
	}

	cfg := gateway.MustReadOrValidateConfig(flags)

	migrate := &Migrator{
		log: log,
		cfg: cfg,
	}

	sqlDB, hostInfo := migrate.setupSqlClient()
	assert.NotNil(t, sqlDB)
	assert.Equal(t, "datalift@0.0.0.0:5432", hostInfo)
}
