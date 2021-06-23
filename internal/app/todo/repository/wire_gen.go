// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package repository

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/middleware/consul"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

// Injectors from wire.go:

func CreateTodoRepository(id string) (TodoRepository, error) {
	client, err := consul.New()
	if err != nil {
		return nil, err
	}
	appConfig := config.NewConfig(id, client)
	rollbarRollbar := rollbar.New(appConfig)
	loggerLogger := logger.NewLogger(rollbarRollbar)
	db, err := mysql.New(appConfig)
	if err != nil {
		return nil, err
	}
	todoRepository := NewMysqlTodoRepository(loggerLogger, db)
	return todoRepository, nil
}

// wire.go:

var testProviderSet = wire.NewSet(logger.ProviderSet, mysql.ProviderSet, config.ProviderSet, consul.ProviderSet, ProviderSet, rollbar.ProviderSet)