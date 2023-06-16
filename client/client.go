package client

import (
	"context"
)

type Client struct {
	ctx    context.Context
	cancel context.CancelFunc
}
