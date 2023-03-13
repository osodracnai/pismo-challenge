package database

import (
	"context"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/opentracing/opentracing-go"
)

type BatchTracer struct{}

func (qt *BatchTracer) ObserveBatch(ctx context.Context, batch gocql.ObservedBatch) {
	if span, _ := opentracing.StartSpanFromContext(ctx, "batch", spanStartOpt{startTime: batch.Start}); span != nil {
		defer span.FinishWithOptions(opentracing.FinishOptions{FinishTime: batch.End})
		// constants
		span.SetTag("span.kind", "client")
		span.SetTag("component", "go-cassandra")
		span.SetTag("db.type", "cassandra")

		// conditional error
		if batch.Err != nil {
			span.SetTag("db.err", batch.Err)
			span.SetTag("error", true)
		}

		// query info
		span.SetTag("db.instance", batch.Keyspace)
		for i, statement := range batch.Statements {
			span.SetTag(fmt.Sprintf("db.statement[%d]", i), statement)
		}
		span.SetTag("db.attempts", batch.Metrics.Attempts)
		span.SetTag("db.total_latency", batch.Metrics.TotalLatency)

		// connection info
		if batch.Host != nil {
			span.SetTag("host", batch.Host.HostID())
			span.SetTag("host.id", batch.Host.HostID())
			span.SetTag("host.dc", batch.Host.DataCenter())
			span.SetTag("peer.hostname", batch.Host.HostnameAndPort())
			span.SetTag("peer.ip", batch.Host.Peer())
			span.SetTag("peer.port", batch.Host.Port())
		}
	}
}
