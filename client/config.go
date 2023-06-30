package client

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"go.datalift.io/datalift/client/defaults"
)

type Config struct {
	HostPort                 string
	AuthToken                string
	DisableTransportSecurity bool
	ConnectionOptions        ConnectionOptions
}

type tokenAuth struct {
	token                    string
	disableTransportSecurity bool
}

func (t tokenAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"Authorization": "Token " + t.token,
	}, nil
}

func (t tokenAuth) RequireTransportSecurity() bool {
	return !t.disableTransportSecurity
}

type ConnectionOptions struct {
	TLSConfig *tls.Config

	DialOptions []grpc.DialOption

	DialTimeout time.Duration

	// Enables keep alive ping from client to the server, which can help detect abruptly closed connections faster.
	EnableKeepAliveCheck bool

	// After a duration of this time if the client doesn't see any activity it
	// pings the server to see if the transport is still alive.
	// If set below 10s, a minimum value of 10s will be used instead.
	KeepAliveTime time.Duration

	// After having pinged for keepalive check, the client waits for a duration
	// of Timeout and if no activity is seen even after that the connection is
	// closed.
	KeepAliveTimeout time.Duration

	// If true, client sends keepalive pings even with no active RPCs. If false,
	// when there are no active RPCs, Time and Timeout will be ignored and no
	// keepalive pings will be sent.
	KeepAlivePermitWithoutStream bool
}

func (c *Config) CheckAndSetDefaults() error {
	// Optional: To set the host:port for this client to connect to.
	// default: localhost:7233
	if c.HostPort == "" {
		c.HostPort = net.JoinHostPort(defaults.DefaultHost, strconv.Itoa(defaults.DefaultPort))
	}

	// Set defaults
	if c.ConnectionOptions.DialTimeout == 0 {
		c.ConnectionOptions.DialTimeout = defaults.DefaultDialTimeout
	}
	if c.ConnectionOptions.KeepAliveTime == 0 {
		c.ConnectionOptions.KeepAliveTime = defaults.DefaultKeepAliveTime
	}
	if c.ConnectionOptions.KeepAliveTimeout == 0 {
		c.ConnectionOptions.KeepAliveTimeout = defaults.DefaultKeepAliveTimeout
	}

	if len(c.AuthToken) == 0 {
		return errors.New("oauth token is required")
	}
	c.ConnectionOptions.DialOptions = append(
		c.ConnectionOptions.DialOptions, grpc.WithPerRPCCredentials(tokenAuth{
			token:                    c.AuthToken,
			disableTransportSecurity: c.DisableTransportSecurity,
		}),
	)

	if c.ConnectionOptions.EnableKeepAliveCheck {
		var kap = keepalive.ClientParameters{
			Time:                c.ConnectionOptions.KeepAliveTime,
			Timeout:             c.ConnectionOptions.KeepAliveTimeout,
			PermitWithoutStream: c.ConnectionOptions.KeepAlivePermitWithoutStream,
		}
		c.ConnectionOptions.DialOptions = append(c.ConnectionOptions.DialOptions, grpc.WithKeepaliveParams(kap))
	}

	return nil
}
