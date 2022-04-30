// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/finance/service"
	repository2 "github.com/tsundata/assistant/internal/app/chatbot/bot/okr/repository"
	service3 "github.com/tsundata/assistant/internal/app/chatbot/bot/okr/service"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/todo/repository"
	service2 "github.com/tsundata/assistant/internal/app/chatbot/bot/todo/service"
	"github.com/tsundata/assistant/internal/app/cron"
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
	rollbarRollbar := rollbar.New(appConfig)
	logger := log.NewZapLogger(rollbarRollbar)
	logLogger := log.NewAppLogger(logger)
	connection, err := rabbitmq.New(appConfig)
	if err != nil {
		return nil, err
	}
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
	messageSvcClient, err := rpcclient.NewMessageClient(rpcClient)
	if err != nil {
		return nil, err
	}
	chatbotSvcClient, err := rpcclient.NewChatbotClient(rpcClient)
	if err != nil {
		return nil, err
	}
	middleSvcClient, err := rpcclient.NewMiddleClient(rpcClient)
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
	locker := global.NewLocker(client)
	idSvcClient, err := rpcclient.NewIdClient(rpcClient)
	if err != nil {
		return nil, err
	}
	globalID := global.NewID(appConfig, idSvcClient)
	conn, err := mysql.New(appConfig)
	if err != nil {
		return nil, err
	}
	todoRepository := repository.NewMysqlTodoRepository(locker, globalID, conn)
	todoSvcServer := service2.NewTodo(bus, logLogger, todoRepository)
	okrRepository := repository2.NewMysqlOkrRepository(locker, globalID, conn)
	okrSvcServer := service3.NewOkr(bus, okrRepository, middleSvcClient)
	componentComponent := component.NewComponent(appConfig, bus, redisClient, logLogger, messageSvcClient, chatbotSvcClient, middleSvcClient, storageSvcClient, userSvcClient, financeSvcServer, todoSvcServer, okrSvcServer)
	ruleBot := rulebot.New(componentComponent)
	application, err := cron.NewApp(appConfig, logLogger, ruleBot)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, log.ProviderSet, rpc.ProviderSet, jaeger.ProviderSet, influx.ProviderSet, redis.ProviderSet, cron.ProviderSet, rollbar.ProviderSet, etcd.ProviderSet, rulebot.ProviderSet, rpcclient.ProviderSet, newrelic.ProviderSet, component.ProviderSet, event.ProviderSet, rabbitmq.ProviderSet, service3.ProviderSet, repository.ProviderSet, service2.ProviderSet, repository2.ProviderSet, service.ProviderSet, global.ProviderSet, mysql.ProviderSet)
