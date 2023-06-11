package meta

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"

	apiv1 "github.com/DataliftHQ/datalift/backend/api/api/v1"
	modulemock "github.com/DataliftHQ/datalift/backend/mock/module"
	"github.com/DataliftHQ/datalift/backend/module/assets"
	"github.com/DataliftHQ/datalift/backend/module/healthcheck"
)

func TestGetAction(t *testing.T) {
	hc, err := healthcheck.New(nil, nil, nil)
	assert.NoError(t, err)

	r := &modulemock.MockRegistrar{Server: grpc.NewServer()}
	err = hc.Register(r)
	assert.NoError(t, err)

	// Register a non-Clutch endpoint with no annotations.
	grpc_health_v1.RegisterHealthServer(r.GRPCServer(), &grpc_health_v1.UnimplementedHealthServer{})

	err = GenerateGRPCMetadata(r.GRPCServer())
	assert.NoError(t, err)

	action := GetAction("/datalift.healthcheck.v1.HealthcheckAPI/Healthcheck")
	assert.Equal(t, apiv1.ActionType_READ, action)

	action = GetAction("/grpc.health.v1.Health/Check")
	assert.Equal(t, apiv1.ActionType_UNSPECIFIED, action)

	action = GetAction("/nonexistent/doesnotexist")
	assert.Equal(t, apiv1.ActionType_UNSPECIFIED, action)
}

func TestAuditDisabled(t *testing.T) {
	r := &modulemock.MockRegistrar{Server: grpc.NewServer()}

	hc, err := healthcheck.New(nil, nil, nil)
	assert.NoError(t, err)
	assert.NoError(t, hc.Register(r))

	a, err := assets.New(nil, nil, nil)
	assert.NoError(t, err)
	assert.NoError(t, a.Register(r))

	assert.NoError(t, GenerateGRPCMetadata(r.GRPCServer()))

	result := IsAuditDisabled("/datalift.healthcheck.v1.HealthcheckAPI/Healthcheck")
	assert.True(t, result)

	result = IsAuditDisabled("/datalift.assets.v1.AssetsAPI/Fetch")
	assert.False(t, result)

	result = IsAuditDisabled("/nonexistent/doesnotexist")
	assert.False(t, result)
}

func TestExtractProtoPatternFieldNames(t *testing.T) {
	t.Parallel()
	tests := []struct {
		id      string
		pattern *apiv1.Pattern
		expect  []string
	}{
		{
			id:      "3 fields",
			pattern: &apiv1.Pattern{Pattern: "{name}/{of}/{fields}"},
			expect:  []string{"name", "of", "fields"},
		},
		{
			id:      "2 fields",
			pattern: &apiv1.Pattern{Pattern: "{name}/{of}"},
			expect:  []string{"name", "of"},
		},
		{
			id:      "1 fields",
			pattern: &apiv1.Pattern{Pattern: "{name}"},
			expect:  []string{"name"},
		},
		{
			id:      "different delimiters",
			pattern: &apiv1.Pattern{Pattern: "{cat}/{meow}-{nom}_{food}--{tasty}"},
			expect:  []string{"cat", "meow", "nom", "food", "tasty"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.id, func(t *testing.T) {
			t.Parallel()

			actual := extractProtoPatternFieldNames(tt.pattern)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestToValue(t *testing.T) {
	type customString string
	type customSlice []bool

	result, err := ToValue(customString("foo"))
	assert.NoError(t, err)
	assert.Equal(t, "foo", result.AsInterface())

	result, err = ToValue(customSlice([]bool{true}))
	assert.NoError(t, err)
	assert.ElementsMatch(t, []bool{true}, result.AsInterface())
}
