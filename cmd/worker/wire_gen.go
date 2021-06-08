// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/worker"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/middleware/consul"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	"github.com/tsundata/assistant/internal/pkg/middleware/influx"
	"github.com/tsundata/assistant/internal/pkg/middleware/jaeger"
	"github.com/tsundata/assistant/internal/pkg/middleware/redis"
	"github.com/tsundata/assistant/internal/pkg/queue"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

// Injectors from wire.go:

func CreateApp() (*app.Application, error) {
	client, err := consul.New()
	if err != nil {
		return nil, err
	}
	appConfig := config.NewConfig(client)
	rollbarRollbar := rollbar.New(appConfig)
	loggerLogger := logger.NewLogger(rollbarRollbar)
	server, err := queue.New(appConfig)
	if err != nil {
		return nil, err
	}
	configuration, err := jaeger.NewConfiguration(appConfig, loggerLogger)
	if err != nil {
		return nil, err
	}
	tracer, err := jaeger.New(configuration)
	if err != nil {
		return nil, err
	}
	clientOptions, err := rpc.NewClientOptions(appConfig, tracer)
	if err != nil {
		return nil, err
	}
	clientv3Client, err := etcd.New(appConfig)
	if err != nil {
		return nil, err
	}
	rpcClient, err := rpc.NewClient(clientOptions, clientv3Client)
	if err != nil {
		return nil, err
	}
	application, err := worker.NewApp(appConfig, loggerLogger, server, rpcClient)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, logger.ProviderSet, http.ProviderSet, rpc.ProviderSet, jaeger.ProviderSet, etcd.ProviderSet, influx.ProviderSet, redis.ProviderSet, worker.ProviderSet, queue.ProviderSet, rollbar.ProviderSet, consul.ProviderSet)
