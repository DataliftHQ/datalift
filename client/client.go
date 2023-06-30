package client

import (
	"compress/gzip"
	"context"
	"fmt"
	"sync/atomic"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	ggzip "google.golang.org/grpc/encoding/gzip"

	applicationpb "go.datalift.io/datalift/api/application/v1"
	"go.datalift.io/datalift/client/metadata"
)

func init() {
	// gzip is used for gRPC auditStream compression. SetLevel changes the
	// compression level, must be called in initialization, and is not thread safe.
	if err := ggzip.SetLevel(gzip.BestSpeed); err != nil {
		panic(err)
	}
}

type Client struct {
	// config contains configuration values for the client.
	config Config

	// conn is a grpc connection to the server.
	conn *grpc.ClientConn

	// grpc is the gRPC client specification for the server.
	grpc serviceClient

	// closedFlag is set to indicate that the connection is closed.
	// It's a pointer to allow the Client struct to be copied.
	closedFlag *int32
}

type serviceClient struct {
	applicationpb.ApplicationAPIClient
}

func New(ctx context.Context, cfg Config) (client *Client, err error) {
	if err = cfg.CheckAndSetDefaults(); err != nil {
		return nil, err
	}

	client = &Client{
		config:     cfg,
		closedFlag: new(int32),
	}
	if err := client.dialGRPC(ctx, cfg.HostPort); err != nil {
		return nil, fmt.Errorf("failed to connect to addr %v as an auth server", cfg.HostPort)
	}
	return client, nil
}

func (c *Client) dialGRPC(ctx context.Context, hostPort string) error {
	dialContext, cancel := context.WithTimeout(ctx, c.config.ConnectionOptions.DialTimeout)
	defer cancel()

	var dialOpts []grpc.DialOption
	dialOpts = append(dialOpts,
		grpc.WithChainUnaryInterceptor(
			metadata.UnaryClientInterceptor,
		),
		grpc.WithChainStreamInterceptor(
			metadata.StreamClientInterceptor,
		),
	)

	// -----------------------------------------------
	// TODO: review
	// cfg

	// Only set transportCredentials if tlsConfig is set. This makes it possible
	// to explicitly provide grpc.WithTransportCredentials(insecure.NewCredentials())
	// in the client's dial options.
	if c.config.ConnectionOptions.TLSConfig != nil {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(c.config.ConnectionOptions.TLSConfig)))
	}
	// -----------------------------------------------

	// must come last, otherwise provided opts may get clobbered by defaults above
	dialOpts = append(dialOpts, c.config.ConnectionOptions.DialOptions...)

	conn, err := grpc.DialContext(dialContext, hostPort, dialOpts...)
	if err != nil {
		return err
	}

	c.conn = conn
	c.grpc = serviceClient{
		ApplicationAPIClient: applicationpb.NewApplicationAPIClient(c.conn),
	}

	return nil
}

// GetConnection returns GRPC connection.
func (c *Client) GetConnection() *grpc.ClientConn {
	return c.conn
}

// Close closes the Client connection to the auth server.
func (c *Client) Close() error {
	if c.setClosed() && c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// isClosed returns whether the client is marked as closed.
//
//nolint:unused
func (c *Client) isClosed() bool {
	return atomic.LoadInt32(c.closedFlag) == 1
}

// setClosed marks the client as closed and returns true if it was open.
func (c *Client) setClosed() bool {
	return atomic.CompareAndSwapInt32(c.closedFlag, 0, 1)
}

func (c *Client) CreateApplication(ctx context.Context, request *applicationpb.CreateApplicationRequest) (*applicationpb.CreateApplicationResponse, error) {
	response, err := c.grpc.CreateApplication(ctx, request)
	if err != nil {
		//return nil, trail.FromGRPC(err)
		return nil, err
	}

	return response, nil
}

func (c *Client) DeleteApplication(ctx context.Context, request *applicationpb.DeleteApplicationRequest) (*applicationpb.DeleteApplicationResponse, error) {
	response, err := c.grpc.DeleteApplication(ctx, request)
	if err != nil {
		//return nil, trail.FromGRPC(err)
		return nil, err
	}

	return response, nil
}

func (c *Client) GetApplication(ctx context.Context, request *applicationpb.GetApplicationRequest) (*applicationpb.GetApplicationResponse, error) {
	response, err := c.grpc.GetApplication(ctx, request)
	if err != nil {
		//return nil, trail.FromGRPC(err)
		return nil, err
	}

	return response, nil
}

func (c *Client) ListApplications(ctx context.Context, request *applicationpb.ListApplicationsRequest) (*applicationpb.ListApplicationsResponse, error) {
	response, err := c.grpc.ListApplications(ctx, request)
	if err != nil {
		//return nil, trail.FromGRPC(err)
		return nil, err
	}

	return response, nil
}

func (c *Client) UpdateApplication(ctx context.Context, request *applicationpb.UpdateApplicationRequest) (*applicationpb.UpdateApplicationResponse, error) {
	response, err := c.grpc.UpdateApplication(ctx, request)
	if err != nil {
		//return nil, trail.FromGRPC(err)
		return nil, err
	}

	return response, nil
}
