package config

import (
	"github.com/sirupsen/logrus"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerPrometheus "github.com/uber/jaeger-lib/metrics/prometheus"
	"io"
	"time"
)

func (j *Jaeger) Config() (io.Closer, error) {

	defcfg := jaegercfg.Configuration{
		ServiceName: j.Name,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	cfg, err := defcfg.FromEnv()
	if err != nil {
		return nil, err
	}

	jLogger := &stdLogger{
		LogError: j.LogError,
		LogInfo:  j.LogInfo,
	}
	jMetricsFactory := jaegerPrometheus.New()

	closer, err := cfg.InitGlobalTracer(
		j.Name,
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		return nil, err
	}
	return closer, nil
}

type stdLogger struct {
	LogError bool
	LogInfo  bool
}

func (l *stdLogger) Error(msg string) {
	if l.LogError {
		logrus.WithFields(logrus.Fields{"jaeger": true}).Error(msg)
	}
}

// Infof logs a message at info priority
func (l *stdLogger) Infof(msg string, args ...interface{}) {
	if l.LogInfo {
		logrus.WithFields(logrus.Fields{"jaeger": true}).Infof(msg, args...)
	}
}
