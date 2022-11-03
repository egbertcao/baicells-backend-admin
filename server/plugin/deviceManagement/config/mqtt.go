package config

import (
	"time"
)

// config holds the configuration
type Mqtt struct {
	Path              string        `mapstructure:"path" json:"path" yaml:"path"`                // ip:port
	ClientID          string        `mapstructure:"clientid" json:"clientid" yaml:"clientid"`    // Client ID to use when connecting to server
	Username          string        `mapstructure:"username" json:"username" yaml:"username"`    // Username to use when connecting to server
	Password          string        `mapstructure:"password" json:"password" yaml:"password"`    // Password to use when connecting to server
	Qos               byte          `mapstructure:"qos" json:"qos" yaml:"qos"`                   // QOS to use when publishing
	KeepAlive         uint16        `mapstructure:"keepalive" json:"keepalive" yaml:"keepalive"` // seconds between keepalive packets
	ConnectRetryDelay time.Duration `mapstructure:"timeout" json:"timeout" yaml:"timeout"`       // Period between connection attempts
}

func MqttDns(m *Mqtt) string {
	return "mqtt://" + m.Username + ":" + m.Password + "@" + m.Path
}
