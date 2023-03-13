package cmd

import (
	"github.com/spf13/pflag"
	"time"
)

func serverFlags(flags *pflag.FlagSet) {
	flags.String("server.listen", "http://0.0.0.0:8080", "HTTP(S) address to listen at.")
}

func cassandraFlags(flags *pflag.FlagSet) {
	flags.StringSlice("cassandra.hosts", []string{"localhost"}, "Cassandra hosts")
	flags.Uint64("cassandra.port", 9042, "Cassandra port")
	flags.String("cassandra.keyspace", "pismo", "Cassandra keyspace")
	flags.String("cassandra.version", "3.0.0", "Cassandra version")
	flags.Duration("cassandra.timeout", 600*time.Millisecond, "Cassandra timeout")
	flags.Duration("cassandra.connect-timeout", 600*time.Millisecond, "Cassandra connect timeout")
	flags.String("cassandra.username", "cassandra", "Cassandra username")
	flags.String("cassandra.password", "cassandra", "Cassandra password")
}

func jaegerFlags(flags *pflag.FlagSet) {
	flags.String("jaeger.name", appName, "Service name used in jaeger")
	flags.Bool("jaeger.log-error", true, "Indicates whether to log jaeger error")
	flags.Bool("jaeger.log-info", false, "Indicates whether to log jaeger information")
}
