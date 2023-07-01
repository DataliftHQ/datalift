package defaults

import "time"

const (
	DefaultHost             = "localhost"
	DefaultPort             = 9443
	DefaultDialTimeout      = 30 * time.Second
	DefaultKeepAliveTime    = 30 * time.Second
	DefaultKeepAliveTimeout = 90 * time.Second
)
