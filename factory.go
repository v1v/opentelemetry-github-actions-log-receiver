package githubactionslogreceiver

import (
	"context"
	"fmt"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

var receiverType = component.MustNewType("githubactionslog")

func createDefaultConfig() component.Config {
	serverConfig := confighttp.NewDefaultServerConfig()
	serverConfig.NetAddr.Endpoint = fmt.Sprintf("localhost:%d", defaultPort)

	return &Config{
		ServerConfig:    serverConfig,
		Path:            defaultPath,
		HealthCheckPath: defaultHealthCheckPath,
		Retry: RetryConfig{
			InitialInterval: defaultRetryInitialInterval,
			MaxInterval:     defaultRetryMaxInterval,
			MaxElapsedTime:  defaultRetryMaxElapsedTime,
		},
		BatchSize: 10000,
	}
}

func createLogsReceiver(
	_ context.Context,
	params receiver.Settings,
	rConf component.Config,
	consumer consumer.Logs,
) (receiver.Logs, error) {
	cfg := rConf.(*Config)
	return newLogsReceiver(cfg, params, consumer)
}

// NewFactory creates a factory for githubactionslogsreceiver.
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		receiverType,
		createDefaultConfig,
		receiver.WithLogs(createLogsReceiver, component.StabilityLevelAlpha),
	)
}
