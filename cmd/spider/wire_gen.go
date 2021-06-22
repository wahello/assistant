// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/spider"
	"github.com/tsundata/assistant/internal/app/spider/rpcclient"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/middleware/consul"
	"github.com/tsundata/assistant/internal/pkg/middleware/influx"
	"github.com/tsundata/assistant/internal/pkg/middleware/jaeger"
	"github.com/tsundata/assistant/internal/pkg/middleware/redis"
	"github.com/tsundata/assistant/internal/pkg/transport/http"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

// Injectors from wire.go:

func CreateApp(id string) (*app.Application, error) {
	client, err := consul.New()
	if err != nil {
		return nil, err
	}
	appConfig := config.NewConfig(id, client)
	redisClient, err := redis.New(appConfig)
	if err != nil {
		return nil, err
	}
	rollbarRollbar := rollbar.New(appConfig)
	loggerLogger := logger.NewLogger(rollbarRollbar)
	configuration, err := jaeger.NewConfiguration(appConfig, loggerLogger)
	if err != nil {
		return nil, err
	}
	tracer, err := jaeger.New(configuration)
	if err != nil {
		return nil, err
	}
	clientOptions, err := rpc.NewClientOptions(tracer)
	if err != nil {
		return nil, err
	}
	rpcClient, err := rpc.NewClient(clientOptions, client, loggerLogger)
	if err != nil {
		return nil, err
	}
	subscribeClient, err := rpcclient.NewSubscribe(rpcClient)
	if err != nil {
		return nil, err
	}
	middleClient, err := rpcclient.NewMiddleClient(rpcClient)
	if err != nil {
		return nil, err
	}
	messageClient, err := rpcclient.NewMessageClient(rpcClient)
	if err != nil {
		return nil, err
	}
	application, err := spider.NewApp(appConfig, redisClient, loggerLogger, subscribeClient, middleClient, messageClient)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, logger.ProviderSet, http.ProviderSet, rpc.ProviderSet, jaeger.ProviderSet, influx.ProviderSet, redis.ProviderSet, spider.ProviderSet, rollbar.ProviderSet, consul.ProviderSet, rpcclient.ProviderSet)
