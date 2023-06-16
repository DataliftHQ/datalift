package aws

import (
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/stretchr/testify/assert"
	"github.com/uber-go/tally/v4"
	"go.uber.org/zap/zaptest"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"testing"

	awsv1 "go.datalift.io/datalift/server/config/service/aws/v1"
)

func TestNew(t *testing.T) {
	regions := []string{"us-east-1", "us-west-2"}

	cfg, _ := anypb.New(&awsv1.Config{
		Regions:                        regions,
		PrimaryAccountAliasDisplayName: "default",
		ClientConfig: &awsv1.ClientConfig{
			Retries: 10,
		},
	})
	log := zaptest.NewLogger(t)
	scope := tally.NewTestScope("", nil)
	s, err := New(cfg, log, scope)
	assert.NoError(t, err)

	// Test conformance to public interface.
	_, ok := s.(Client)
	assert.True(t, ok)

	// Test private interface.
	c, ok := s.(*client)
	assert.True(t, ok)

	assert.NotNil(t, c.log)
	assert.NotNil(t, c.scope)

	assert.Len(t, c.accounts["default"].clients, len(regions))
	addedRegions := make([]string, 0, len(regions))

	for key, rc := range c.accounts[c.currentAccountAlias].clients {
		addedRegions = append(addedRegions, key)
		assert.Equal(t, key, rc.region)
	}
	assert.ElementsMatch(t, addedRegions, regions)
}

func TestNewWithWrongConfigType(t *testing.T) {
	_, err := New(&any.Any{TypeUrl: "foo"}, nil, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "mismatched message type")
}

func TestConfigureAdditionalAccountClient(t *testing.T) {
	cfg, _ := anypb.New(&awsv1.Config{
		Regions:                        []string{"us-east-1"},
		PrimaryAccountAliasDisplayName: "default",
		AdditionalAccounts: []*awsv1.AWSAccount{
			{
				Alias:         "dev",
				AccountNumber: "123",
				IamRole:       "iam-dev",
				Regions:       []string{"us-west-1"},
			},
			{
				Alias:         "staging",
				AccountNumber: "456",
				IamRole:       "iam-staging",
				Regions:       []string{"us-west-2"},
			},
		},
	})
	log := zaptest.NewLogger(t)
	scope := tally.NewTestScope("", nil)
	s, err := New(cfg, log, scope)
	assert.NoError(t, err)

	c, ok := s.(*client)
	assert.True(t, ok)

	expctedAccounts := []string{"default", "dev", "staging"}
	for _, expect := range expctedAccounts {
		val, ok := c.accounts[expect]
		assert.True(t, ok)
		assert.NotNil(t, val)
	}
}

func TestGetAccountRegionClient(t *testing.T) {
	c := &client{
		currentAccountAlias: "default",
		accounts: map[string]*accountClients{
			"default": {
				clients: map[string]*regionalClient{
					"us-east-1":  nil,
					"us-west-2":  nil,
					"us-north-5": nil,
				},
			},
		},
	}

	_, err := c.getAccountRegionClient("default", "us-east-1")
	assert.NoError(t, err)

	_, err = c.getAccountRegionClient("aws://", "us-west-3")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")

	_, err = c.getAccountRegionClient("default", "us-west-3")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no client found")
}

func TestRegions(t *testing.T) {
	c := &client{
		currentAccountAlias: "default",
		accounts: map[string]*accountClients{
			"default": {
				clients: map[string]*regionalClient{
					"us-east-1":  nil,
					"us-west-2":  nil,
					"us-north-5": nil,
				},
			},
		},
	}

	regions := c.Regions()
	assert.ElementsMatch(t, regions, []string{"us-east-1", "us-west-2", "us-north-5"})
}

func TestGetAccountsInRegion(t *testing.T) {
	c := &client{
		currentAccountAlias: "default",
		accounts: map[string]*accountClients{
			"default": {
				alias:   "default",
				regions: []string{"us-east-1", "us-west-2", "us-north-5"},
			},
			"staging": {
				alias:   "staging",
				regions: []string{"us-east-1"},
			},
			"prod": {
				alias:   "prod",
				regions: []string{"us-east-2"},
			},
			"testing": {
				alias:   "testing",
				regions: []string{"us-west-2"},
			},
		},
	}

	assert.ElementsMatch(t, []string{"default", "staging"}, c.GetAccountsInRegion("us-east-1"))
	assert.ElementsMatch(t, []string{"prod"}, c.GetAccountsInRegion("us-east-2"))
	assert.ElementsMatch(t, []string{"default", "testing"}, c.GetAccountsInRegion("us-west-2"))
	assert.ElementsMatch(t, []string{"default"}, c.GetAccountsInRegion("us-north-5"))
}

func TestDuplicateRegions(t *testing.T) {
	c := &client{
		currentAccountAlias: "default",
		accounts: map[string]*accountClients{
			"default": {
				clients: map[string]*regionalClient{
					"us-east-1":  nil,
					"us-west-2":  nil,
					"us-north-5": nil,
				},
			},
			"staging": {
				clients: map[string]*regionalClient{
					"us-east-1":  nil,
					"us-west-2":  nil,
					"us-north-1": nil,
					"us-north-5": nil,
				},
			},
		},
	}

	regions := c.Regions()
	assert.ElementsMatch(t, regions, []string{"us-east-1", "us-west-2", "us-north-1", "us-north-5"})
}

func TestZoneToRegion(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{
			input:  "us-east-1a",
			expect: "us-east-1",
		},
		{
			input:  "",
			expect: "UNKNOWN",
		},
	}

	for _, test := range tests {
		output := zoneToRegion(test.input)
		assert.Equal(t, test.expect, output)
	}
}

func TestErrorIntercept(t *testing.T) {
	c := &client{}
	{
		origErr := newResponseError(400, &smithy.GenericAPIError{Code: "whoopsie", Message: "bad"})
		err := c.InterceptError(origErr)
		_, ok := status.FromError(err)
		assert.True(t, ok)
	}
	{
		origErr := errors.New("foo")
		err := c.InterceptError(origErr)
		assert.Equal(t, origErr, err)
	}
}

func TestRegionalClient(t *testing.T) {
	c := &s3.Client{}

	r := &regionalClient{s3: c}

	assert.Equal(t, c, r.S3())
}

func TestRegionalClientConfig(t *testing.T) {
	c := &aws.Config{Region: "us-west-1"}

	r := &regionalClient{regionCfg: c}

	assert.Equal(t, c, r.Config())
}
