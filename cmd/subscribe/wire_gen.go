// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/subscribe"
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

func CreateApp(cf string) (*app.Application, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	options, err := subscribe.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	rollbarOptions, err := rollbar.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	rollbarRollbar := rollbar.New(rollbarOptions)
	loggerLogger := logger.NewLogger(rollbarRollbar)
	rpcOptions, err := rpc.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	configuration, err := jaeger.NewConfiguration(viper, loggerLogger)
	if err != nil {
		return nil, err
	}
	tracer, err := jaeger.New(configuration)
	if err != nil {
		return nil, err
	}
	etcdOptions, err := etcd.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	client, err := etcd.New(etcdOptions)
	if err != nil {
		return nil, err
	}
	influxOptions, err := influx.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	influxdb2Client, err := influx.New(influxOptions)
	if err != nil {
		return nil, err
	}
	server, err := rpc.NewServer(rpcOptions, loggerLogger, tracer, client, influxdb2Client)
	if err != nil {
		return nil, err
	}
	application, err := subscribe.NewApp(options, loggerLogger, server, client)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, logger.ProviderSet, http.ProviderSet, rpc.ProviderSet, jaeger.ProviderSet, etcd.ProviderSet, influx.ProviderSet, redis.ProviderSet, subscribe.ProviderSet, database.ProviderSet, rollbar.ProviderSet)
