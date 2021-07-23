// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package config

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
)

// Injectors from wire.go:

func CreateAppConfig(id string) (*AppConfig, error) {
	client, err := etcd.New()
	if err != nil {
		return nil, err
	}
	appConfig := NewConfig(id, client)
	return appConfig, nil
}

// wire.go:

var testProviderSet = wire.NewSet(etcd.ProviderSet, ProviderSet)
