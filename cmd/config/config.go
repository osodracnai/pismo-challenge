package config

import "time"

type (
	Global struct {
		Cassandra Cassandra
		Server    Server
		Jaeger    Jaeger
	}
	Server struct {
		Listen string `mapstructure:"listen"`
		Debug  bool   `mapstructure:"debug"`
	}
	Jaeger struct {
		Name     string
		LogError bool `mapstructure:"log-error"`
		LogInfo  bool `mapstructure:"log-info"`
	}
	Cassandra struct {
		Hosts          []string
		Port           int
		Keyspace       string
		Version        string
		Timeout        time.Duration
		ConnectTimeout time.Duration `mapstructure:"connect-timeout"`
		Username       string
		Password       string
	}
)
