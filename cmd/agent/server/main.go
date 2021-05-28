package main

import (
	"github.com/tsundata/assistant/internal/app/agent"
	"github.com/tsundata/assistant/internal/app/agent/broker"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/influx"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

func CreateApp() (*app.Application, error) {
	appConfig := config.NewConfig()

	r := rollbar.New(appConfig)

	log := logger.NewLogger(r)

	i, err := influx.New(appConfig)
	if err != nil {
		return nil, err
	}

	br, err := broker.NewBroker(appConfig, log, i)
	if err != nil {
		return nil, err
	}

	b, err := broker.NewServer(br)
	if err != nil {
		return nil, err
	}

	application, err := agent.NewApp(appConfig, log, b)
	if err != nil {
		return nil, err
	}
	return application, nil
}

func main() {
	a, err := CreateApp()
	if err != nil {
		panic(err)
	}

	if err := a.Start(); err != nil {
		panic(err)
	}

	a.AwaitSignal()
}
