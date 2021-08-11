// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/chatbot"
	"github.com/tsundata/assistant/internal/app/chatbot/rpcclient"
	"github.com/tsundata/assistant/internal/app/chatbot/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	"github.com/tsundata/assistant/internal/pkg/middleware/influx"
	"github.com/tsundata/assistant/internal/pkg/middleware/jaeger"
	"github.com/tsundata/assistant/internal/pkg/middleware/nats"
	"github.com/tsundata/assistant/internal/pkg/middleware/redis"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
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
	middleSvcClient, err := rpcclient.NewMiddleClient(rpcClient)
	if err != nil {
		return nil, err
	}
	todoSvcClient, err := rpcclient.NewTodoClient(rpcClient)
	if err != nil {
		return nil, err
	}
	redisClient, err := redis.New(appConfig, newrelicApp)
	if err != nil {
		return nil, err
	}
	messageSvcClient, err := rpcclient.NewMessageClient(rpcClient)
	if err != nil {
		return nil, err
	}
	workflowSvcClient, err := rpcclient.NewWorkflowClient(rpcClient)
	if err != nil {
		return nil, err
	}
	storageSvcClient, err := rpcclient.NewStorageClient(rpcClient)
	if err != nil {
		return nil, err
	}
	userSvcClient, err := rpcclient.NewUserClient(rpcClient)
	if err != nil {
		return nil, err
	}
	nlpSvcClient, err := rpcclient.NewNLPClient(rpcClient)
	if err != nil {
		return nil, err
	}
	orgSvcClient, err := rpcclient.NewOrgClient(rpcClient)
	if err != nil {
		return nil, err
	}
	iComponent := rulebot.NewComponent(appConfig, redisClient, logLogger, messageSvcClient, middleSvcClient, workflowSvcClient, storageSvcClient, todoSvcClient, userSvcClient, nlpSvcClient, orgSvcClient)
	ruleBot := rulebot.New(iComponent)
	serviceChatbot := service.NewChatbot(logLogger, middleSvcClient, todoSvcClient, ruleBot)
	initServer := service.CreateInitServerFn(serviceChatbot)
	server, err := rpc.NewServer(appConfig, logger, logLogger, initServer, tracer, redisClient, newrelicApp)
	if err != nil {
		return nil, err
	}
	application, err := chatbot.NewApp(appConfig, bus, logLogger, server, middleSvcClient, todoSvcClient, userSvcClient)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, log.ProviderSet, rpc.ProviderSet, jaeger.ProviderSet, influx.ProviderSet, redis.ProviderSet, chatbot.ProviderSet, rollbar.ProviderSet, etcd.ProviderSet, service.ProviderSet, rpcclient.ProviderSet, rulebot.ProviderSet, event.ProviderSet, nats.ProviderSet, newrelic.ProviderSet)
