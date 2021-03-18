package main

import (
	"flag"
	"github.com/tsundata/assistant/internal/app/worker"
	"github.com/tsundata/assistant/internal/app/worker/rpcclients"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/etcd"
	"github.com/tsundata/assistant/internal/pkg/jaeger"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/machinery"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

func CreateApp(name, cf string) (*app.Application, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	log := logger.NewLogger()

	rollbarOptions, err := rollbar.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	rollbar.Config(rollbarOptions)

	t, err := jaeger.NewConfiguration(viper, log)
	if err != nil {
		return nil, err
	}
	j, err := jaeger.New(t)
	if err != nil {
		return nil, err
	}

	etcdOption, err := etcd.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	e, err := etcd.New(etcdOption)
	if err != nil {
		return nil, err
	}

	machineryOption, err := machinery.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	ms, err := machinery.New(machineryOption)
	if err != nil {
		return nil, err
	}

	clientOptions, err := rpc.NewClientOptions(viper, j)
	if err != nil {
		return nil, err
	}
	client, err := rpc.NewClient(clientOptions, e)
	if err != nil {
		return nil, err
	}
	msgClient, err := rpcclients.NewMessageClient(client)
	if err != nil {
		return nil, err
	}
	wfClient, err := rpcclients.NewWorkflowClient(client)
	if err != nil {
		return nil, err
	}

	application, err := worker.NewApp(name, log, ms, msgClient, wfClient)
	if err != nil {
		return nil, err
	}
	return application, nil
}

var appName = flag.String("n", "appName", "set app name")
var configFile = flag.String("f", "worker.yml", "set config file which will loading")

func main() {
	flag.Parse()

	a, err := CreateApp(*appName, *configFile)
	if err != nil {
		panic(err)
	}

	if err := a.Start(); err != nil {
		panic(err)
	}

	a.AwaitSignal()
}