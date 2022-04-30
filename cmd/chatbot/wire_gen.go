// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/chatbot"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/finance/service"
	repository3 "github.com/tsundata/assistant/internal/app/chatbot/bot/okr/repository"
	service3 "github.com/tsundata/assistant/internal/app/chatbot/bot/okr/service"
	repository2 "github.com/tsundata/assistant/internal/app/chatbot/bot/todo/repository"
	service2 "github.com/tsundata/assistant/internal/app/chatbot/bot/todo/service"
	"github.com/tsundata/assistant/internal/app/chatbot/repository"
	service4 "github.com/tsundata/assistant/internal/app/chatbot/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	"github.com/tsundata/assistant/internal/pkg/middleware/influx"
	"github.com/tsundata/assistant/internal/pkg/middleware/jaeger"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/middleware/rabbitmq"
	"github.com/tsundata/assistant/internal/pkg/middleware/redis"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
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
	bus := event.NewRabbitmqBus(connection, logLogger)
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
	locker := global.NewLocker(client)
	conn, err := mysql.New(appConfig)
	if err != nil {
		return nil, err
	}
	chatbotRepository := repository.NewMysqlChatbotRepository(globalID, locker, conn)
	messageSvcClient, err := rpcclient.NewMessageClient(rpcClient)
	if err != nil {
		return nil, err
	}
	middleSvcClient, err := rpcclient.NewMiddleClient(rpcClient)
	if err != nil {
		return nil, err
	}
	chatbotSvcClient, err := rpcclient.NewChatbotClient(rpcClient)
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
	financeSvcServer := service.NewFinance()
	todoRepository := repository2.NewMysqlTodoRepository(locker, globalID, conn)
	todoSvcServer := service2.NewTodo(bus, logLogger, todoRepository)
	okrRepository := repository3.NewMysqlOkrRepository(locker, globalID, conn)
	okrSvcServer := service3.NewOkr(bus, okrRepository, middleSvcClient)
	componentComponent := component.NewComponent(appConfig, bus, redisClient, logLogger, messageSvcClient, chatbotSvcClient, middleSvcClient, storageSvcClient, userSvcClient, financeSvcServer, todoSvcServer, okrSvcServer)
	ruleBot := rulebot.New(componentComponent)
	serviceChatbot := service4.NewChatbot(logLogger, bus, redisClient, chatbotRepository, messageSvcClient, middleSvcClient, ruleBot, componentComponent)
	initServer := service4.CreateInitServerFn(serviceChatbot)
	server, err := rpc.NewServer(appConfig, logger, logLogger, initServer, tracer, redisClient, newrelicApp)
	if err != nil {
		return nil, err
	}
	application, err := chatbot.NewApp(appConfig, bus, redisClient, logLogger, server, messageSvcClient, middleSvcClient, chatbotRepository, ruleBot, componentComponent)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, log.ProviderSet, rpc.ProviderSet, jaeger.ProviderSet, influx.ProviderSet, redis.ProviderSet, chatbot.ProviderSet, rollbar.ProviderSet, etcd.ProviderSet, service4.ProviderSet, rpcclient.ProviderSet, rulebot.ProviderSet, event.ProviderSet, newrelic.ProviderSet, mysql.ProviderSet, repository.ProviderSet, global.ProviderSet, rabbitmq.ProviderSet, component.ProviderSet, service3.ProviderSet, repository2.ProviderSet, service2.ProviderSet, repository3.ProviderSet, service.ProviderSet)
