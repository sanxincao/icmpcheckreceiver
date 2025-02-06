package icmpcheckreceiver

import (
	"time"
)

type Config struct {
	Interval time.Duration `mapstructure:"interval"`
	Targets  []Target      `mapstructure:"targets"`
}

type Target struct {
	Target      string         `mapstructure:"target"`
	PingCount   *int           `mapstructure:"ping_count"`
	PingTimeout *time.Duration `mapstructure:"ping_timeout"`
}
