package client

import (
	"context"
)

// TODO: build go client package that can be used by cli, terraform provider, and more!

type Client struct {
	ctx    context.Context
	cancel context.CancelFunc
}
