package aws

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"io"
	"net/http"

	awsv1 "github.com/DataliftHQ/datalift/backend/config/service/aws/v1"
	"github.com/DataliftHQ/datalift/backend/service"
)

const (
	Name = "datalift.service.aws"
)

func New(cfg *anypb.Any, logger *zap.Logger, scope tally.Scope) (service.Service, error) {
	ac := &awsv1.Config{}
	err := cfg.UnmarshalTo(ac)
	if err != nil {
		return nil, err
	}

	// aws_config_profile_name is not currently implemented
	// if this is set will error out to let the user know what they are trying to do will not work
	if ac.AwsConfigProfileName != "" {
		return nil, errors.New("AWS config field [aws_config_profile_name] is not implemented")
	}

	accountAlias := ac.PrimaryAccountAliasDisplayName
	if ac.PrimaryAccountAliasDisplayName == "" {
		accountAlias = "default"
	}

	c := &client{
		accounts:            make(map[string]*accountClients),
		currentAccountAlias: accountAlias,
		log:                 logger,
		scope:               scope,
	}

	clientRetries := 0
	if ac.ClientConfig != nil && ac.ClientConfig.Retries >= 0 {
		clientRetries = int(ac.ClientConfig.Retries)
	}

	awsHTTPClient := &http.Client{}
	awsClientCommonOptions := []func(*config.LoadOptions) error{
		config.WithHTTPClient(awsHTTPClient),
		config.WithRetryer(func() aws.Retryer {
			customRetryer := retry.NewStandard(func(so *retry.StandardOptions) {
				so.MaxAttempts = clientRetries
			})
			return customRetryer
		}),
	}

	for _, region := range ac.Regions {
		regionCfg, err := config.LoadDefaultConfig(context.TODO(),
			append(awsClientCommonOptions, config.WithRegion(region))...,
		)
		if err != nil {
			return nil, err
		}

		c.createRegionalClients(c.currentAccountAlias, region, ac.Regions, regionCfg)
	}

	if err := c.configureAdditionalAccountClient(ac.AdditionalAccounts, awsHTTPClient, awsClientCommonOptions); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *client) configureAdditionalAccountClient(accounts []*awsv1.AWSAccount, awsHTTPClient *http.Client, awsClientOptions []func(*config.LoadOptions) error) error {
	for _, account := range accounts {
		accountRoleARN := fmt.Sprintf("arn:aws:iam::%s:role/%s", account.AccountNumber, account.IamRole)
		// For doing STS calls it does not matter which region client we are using, as they are not bounded by region
		// we choose just the first region client
		stsClient := c.accounts[c.currentAccountAlias].clients[c.accounts[c.currentAccountAlias].regions[0]].sts
		assumeRoleProvider := stscreds.NewAssumeRoleProvider(stsClient, accountRoleARN)
		credsCache := aws.NewCredentialsCache(assumeRoleProvider)

		for _, region := range account.Regions {
			regionCfg, err := config.LoadDefaultConfig(context.TODO(),
				append(awsClientOptions, config.WithRegion(region))...,
			)
			if err != nil {
				return err
			}

			regionCfg.Credentials = credsCache

			c.createRegionalClients(account.Alias, region, account.Regions, regionCfg)
		}
	}

	return nil
}

func (c *client) createRegionalClients(accountAlias, region string, regions []string, regionCfg aws.Config) {
	if _, ok := c.accounts[accountAlias]; !ok {
		c.accounts[accountAlias] = &accountClients{
			alias:   accountAlias,
			regions: regions,
			clients: map[string]*regionalClient{},
		}
	}

	c.accounts[accountAlias].clients[region] = &regionalClient{
		region:    region,
		regionCfg: &regionCfg,

		iam: iam.NewFromConfig(regionCfg),
		s3:  s3.NewFromConfig(regionCfg),
		sts: sts.NewFromConfig(regionCfg),
	}
}

type Client interface {
	S3GetBucketPolicy(ctx context.Context, account, region, bucket, accountID string) (*s3.GetBucketPolicyOutput, error)
	S3StreamingGet(ctx context.Context, account, region, bucket, key string) (io.ReadCloser, error)

	GetCallerIdentity(ctx context.Context, account, region string) (*sts.GetCallerIdentityOutput, error)

	SimulateCustomPolicy(ctx context.Context, account, region string, customPolicySimulatorParams *iam.SimulateCustomPolicyInput) (*iam.SimulateCustomPolicyOutput, error)

	Accounts() []string
	AccountsAndRegions() map[string][]string
	GetAccountsInRegion(region string) []string
	GetPrimaryAccountAlias() string
	Regions() []string

	GetDirectClient(account string, region string) (DirectClient, error)
}

// DirectClient gives access to the underlying AWS clients from the Golang SDK.
// This allows arbitrary feature development on top of AWS from other services and modules without having to
// contribute to the upstream interface. Using these clients will make mocking extremely difficult since it returns the
// AWS SDK's struct types and not an interface that can be substituted for. It is recommended following initial
// development of a feature that you add the calls to a service interface so they can be tested more easily.
type DirectClient interface {
	Config() *aws.Config
	IAM() *iam.Client
	S3() *s3.Client
	STS() *sts.Client
}

type client struct {
	accounts            map[string]*accountClients
	currentAccountAlias string
	log                 *zap.Logger
	scope               tally.Scope
}

type regionalClient struct {
	region    string
	regionCfg *aws.Config

	iam iamClient
	s3  s3Client
	sts stsClient
}

func (r *regionalClient) Config() *aws.Config {
	return r.regionCfg
}

func (r *regionalClient) IAM() *iam.Client {
	return r.iam.(*iam.Client)
}

func (r *regionalClient) S3() *s3.Client {
	return r.s3.(*s3.Client)
}

func (r *regionalClient) STS() *sts.Client {
	return r.sts.(*sts.Client)
}

type accountClients struct {
	alias   string
	regions []string

	clients map[string]*regionalClient
}

func (c *client) GetDirectClient(account string, region string) (DirectClient, error) {
	return c.getAccountRegionClient(account, region)
}

// Implement the interface provided by errorintercept, so errors are caught at middleware and converted to gRPC status.
func (c *client) InterceptError(e error) error {
	return ConvertError(e)
}

func (c *client) getAccountRegionClient(account, region string) (*regionalClient, error) {
	accountClients, ok := c.accounts[account]
	if !ok || accountClients == nil {
		return nil, status.Errorf(codes.NotFound, "account %s not found", account)
	}
	cl, ok := accountClients.clients[region]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "no client found for account '%s' in region '%s'", account, region)
	}
	return cl, nil
}

func (c *client) Regions() []string {
	uniqueRegions := map[string]bool{}

	for _, account := range c.accounts {
		for region := range account.clients {
			uniqueRegions[region] = true
		}
	}

	regions := make([]string, len(uniqueRegions))
	i := 0
	for region := range uniqueRegions {
		regions[i] = region
		i++
	}

	return regions
}

func (c *client) Accounts() []string {
	accounts := []string{}
	for account := range c.accounts {
		accounts = append(accounts, account)
	}
	return accounts
}

func (c *client) AccountsAndRegions() map[string][]string {
	ar := make(map[string][]string)
	for name, account := range c.accounts {
		ar[name] = account.regions
	}
	return ar
}

// Get all accounts that exist in a specific region
func (c *client) GetAccountsInRegion(region string) []string {
	accounts := []string{}
	for _, a := range c.accounts {
		for _, r := range a.regions {
			if r == region {
				accounts = append(accounts, a.alias)
			}
		}
	}

	return accounts
}

// Shave off the trailing zone identifier to get the region
func zoneToRegion(zone string) string {
	if zone == "" {
		return "UNKNOWN"
	}
	return zone[:len(zone)-1]
}

func (c *client) GetPrimaryAccountAlias() string {
	return c.currentAccountAlias
}
