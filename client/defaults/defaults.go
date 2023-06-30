package defaults

import "time"

const (
	DefaultHost             = "localhost"
	DefaultPort             = 8080
	DefaultDialTimeout      = 30 * time.Second
	DefaultKeepAliveTime    = 30 * time.Second
	DefaultKeepAliveTimeout = 90 * time.Second
)
