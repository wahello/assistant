// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/workflow"
	"github.com/tsundata/assistant/internal/app/workflow/repository"
	"github.com/tsundata/assistant/internal/app/workflow/rpcclient"
	"github.com/tsundata/assistant/internal/app/workflow/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	"github.com/tsundata/assistant/internal/pkg/middleware/influx"
	"github.com/tsundata/assistant/internal/pkg/middleware/jaeger"
	"github.com/tsundata/assistant/internal/pkg/middleware/nats"
	"github.com/tsundata/assistant/internal/pkg/middleware/redis"
	"github.com/tsundata/assistant/internal/pkg/middleware/rqlite"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

// Injectors from wire.go:

func CreateApp(id string) (*app.Application, error) {
	client, err := etcd.New()
	if err != nil {
		return nil, err
	}
	appConfig := config.NewConfig(id, client)
	conn, err := nats.New(appConfig)
	if err != nil {
		return nil, err
	}
	rollbarRollbar := rollbar.New(appConfig)
	logger := log.NewZapLogger(rollbarRollbar)
	newrelicApp, err := newrelic.New(appConfig, logger)
	if err != nil {
		return nil, err
	}
	bus := event.NewNatsBus(conn, newrelicApp)
	logLogger := log.NewAppLogger(logger)
	redisClient, err := redis.New(appConfig, newrelicApp)
	if err != nil {
		return nil, err
	}
	rqliteConn, err := rqlite.New(appConfig, newrelicApp, logLogger)
	if err != nil {
		return nil, err
	}
	workflowRepository := repository.NewRqliteWorkflowRepository(rqliteConn)
	configuration, err := jaeger.NewConfiguration(appConfig, logLogger)
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
	rpcClient, err := rpc.NewClient(clientOptions, appConfig, logLogger)
	if err != nil {
		return nil, err
	}
	messageSvcClient, err := rpcclient.NewMessageClient(rpcClient)
	if err != nil {
		return nil, err
	}
	middleSvcClient, err := rpcclient.NewMiddleClient(rpcClient)
	if err != nil {
		return nil, err
	}
	serviceWorkflow := service.NewWorkflow(bus, redisClient, workflowRepository, messageSvcClient, middleSvcClient, logLogger)
	initServer := service.CreateInitServerFn(serviceWorkflow)
	server, err := rpc.NewServer(appConfig, logger, logLogger, initServer, tracer, redisClient, newrelicApp)
	if err != nil {
		return nil, err
	}
	application, err := workflow.NewApp(appConfig, bus, logLogger, server, redisClient, middleSvcClient, messageSvcClient)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, log.ProviderSet, rpc.ProviderSet, jaeger.ProviderSet, influx.ProviderSet, redis.ProviderSet, workflow.ProviderSet, rollbar.ProviderSet, repository.ProviderSet, etcd.ProviderSet, event.ProviderSet, nats.ProviderSet, service.ProviderSet, rpcclient.ProviderSet, newrelic.ProviderSet, rqlite.ProviderSet)
