// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/user"
	"github.com/tsundata/assistant/internal/app/user/repository"
	"github.com/tsundata/assistant/internal/app/user/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	"github.com/tsundata/assistant/internal/pkg/middleware/influx"
	"github.com/tsundata/assistant/internal/pkg/middleware/jaeger"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/middleware/nats"
	"github.com/tsundata/assistant/internal/pkg/middleware/rabbitmq"
	"github.com/tsundata/assistant/internal/pkg/middleware/redis"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/rpcclient"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

func CreateApp(id string) (*app.Application, error) {
	client, err := etcd.New()
	if err != nil {
		return nil, err
	}
	appConfig := config.NewConfig(id, client)
	connection, err := rabbitmq.New(appConfig)
	if err != nil {
		return nil, err
	}
	rollbarRollbar := rollbar.New(appConfig)
	logger := log.NewZapLogger(rollbarRollbar)
	logLogger := log.NewAppLogger(logger)
	bus := event.NewNatsBus(connection, logLogger)
	newrelicApp, err := newrelic.New(appConfig, logger)
	if err != nil {
		return nil, err
	}
	redisClient, err := redis.New(appConfig, newrelicApp)
	if err != nil {
		return nil, err
	}
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
	idSvcClient, err := rpcclient.NewIdClient(rpcClient)
	if err != nil {
		return nil, err
	}
	globalID := global.NewID(appConfig, idSvcClient)
	conn, err := mysql.New(appConfig)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewMysqlUserRepository(logLogger, globalID, conn)
	serviceUser := service.NewUser(appConfig, logLogger, redisClient, userRepository)
	initServer := service.CreateInitServerFn(serviceUser)
	server, err := rpc.NewServer(appConfig, logger, logLogger, initServer, tracer, redisClient, newrelicApp)
	if err != nil {
		return nil, err
	}
	nlpSvcClient, err := rpcclient.NewNLPClient(rpcClient)
	if err != nil {
		return nil, err
	}
	application, err := user.NewApp(appConfig, bus, logLogger, server, userRepository, nlpSvcClient)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, log.ProviderSet, rpc.ProviderSet, jaeger.ProviderSet, influx.ProviderSet, redis.ProviderSet, user.ProviderSet, rollbar.ProviderSet, etcd.ProviderSet, repository.ProviderSet, event.ProviderSet, nats.ProviderSet, service.ProviderSet, rpcclient.ProviderSet, newrelic.ProviderSet, mysql.ProviderSet, global.ProviderSet, rabbitmq.ProviderSet)
