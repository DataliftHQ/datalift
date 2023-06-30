package migrator

import (
	"github.com/stretchr/testify/assert"
	"go.datalift.io/datalift/internal/gateway"
	"go.uber.org/zap/zaptest"
	"os"
	"path"
	"testing"
)

func TestSetupSqlClient(t *testing.T) {
	log := zaptest.NewLogger(t)

	pwd, err := os.Getwd()
	assert.NoError(t, err)

	flags := &gateway.Flags{
		ConfigPath: path.Join(pwd, "testdata/datalift-test.yaml"),
	}

	cfg := gateway.MustReadOrValidateConfig(flags)

	migrator := &Migrator{
		log:    log,
		config: cfg,
	}

	sqlDB, hostInfo := migrator.setupSqlClient()
	assert.NotNil(t, sqlDB)
	assert.Equal(t, "datalift@0.0.0.0:5432", hostInfo)
}
