// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/worker"
	"github.com/tsundata/assistant/internal/app/worker/rpcclients"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/etcd"
	"github.com/tsundata/assistant/internal/pkg/influx"
	"github.com/tsundata/assistant/internal/pkg/jaeger"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/machinery"
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
	options, err := rollbar.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	rollbarRollbar := rollbar.New(options)
	loggerLogger := logger.NewLogger(rollbarRollbar)
	machineryOptions, err := machinery.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	server, err := machinery.New(machineryOptions)
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
	clientOptions, err := rpc.NewClientOptions(viper, tracer)
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
	rpcClient, err := rpc.NewClient(clientOptions, client)
	if err != nil {
		return nil, err
	}
	messageClient, err := rpcclients.NewMessageClient(rpcClient)
	if err != nil {
		return nil, err
	}
	workflowClient, err := rpcclients.NewWorkflowClient(rpcClient)
	if err != nil {
		return nil, err
	}
	application, err := worker.NewApp(loggerLogger, server, messageClient, workflowClient)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, logger.ProviderSet, http.ProviderSet, rpc.ProviderSet, jaeger.ProviderSet, etcd.ProviderSet, influx.ProviderSet, redis.ProviderSet, worker.ProviderSet, rpcclients.ProviderSet, machinery.ProviderSet, rollbar.ProviderSet)
