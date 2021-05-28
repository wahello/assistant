// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/storage"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/database"
	"github.com/tsundata/assistant/internal/pkg/etcd"
	"github.com/tsundata/assistant/internal/pkg/influx"
	"github.com/tsundata/assistant/internal/pkg/jaeger"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/redis"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

// Injectors from wire.go:

func CreateApp() (*app.Application, error) {
	appConfig := config.NewConfig()
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
	client, err := etcd.New(appConfig)
	if err != nil {
		return nil, err
	}
	influxdb2Client, err := influx.New(appConfig)
	if err != nil {
		return nil, err
	}
	redisClient, err := redis.New(appConfig)
	if err != nil {
		return nil, err
	}
	server, err := rpc.NewServer(appConfig, loggerLogger, tracer, client, influxdb2Client, redisClient)
	if err != nil {
		return nil, err
	}
	db, err := database.New(appConfig)
	if err != nil {
		return nil, err
	}
	application, err := workflow.NewApp(appConfig, loggerLogger, server, client, db, redisClient)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, logger.ProviderSet, http.ProviderSet, rpc.ProviderSet, jaeger.ProviderSet, etcd.ProviderSet, influx.ProviderSet, redis.ProviderSet, workflow.ProviderSet, database.ProviderSet, rollbar.ProviderSet)