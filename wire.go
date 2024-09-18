//go:build wireinject
// +build wireinject

package main

import (
	"be-golang-chapter-26/config"
	"be-golang-chapter-26/greeter"
	"be-golang-chapter-26/notification"
	"be-golang-chapter-26/service"
	"be-golang-chapter-26/storage"

	"github.com/google/wire"
)

// func InitializMyService() (*service.Service, error) {
// 	wire.Build(greeter.NewGreeter, service.NewService)
// 	return nil, nil
// }

// func InitializMyService(name string) (*service.Service, error) {
// 	wire.Build(greeter.NewGreeter, service.NewService)
// 	return nil, nil
// }

// func InitializeServiceConfig() (*service.ServiceConfig, error) {
// 	wire.Build(config.NewConfig, config.NewConfigAlternative, service.NewServiceConfig)
// 	return nil, nil
// }

var myservice = wire.NewSet(
	greeter.NewGreeter,
	service.NewService,
)

func InitializMyService(name string) (*service.Service, error) {
	wire.Build(myservice)
	return nil, nil
}

var serviceConfig = wire.NewSet(
	config.NewConfig,
	config.NewConfigAlternative,
	service.NewServiceConfig,
)

func InitializeServiceConfig() (*service.ServiceConfig, error) {
	wire.Build(serviceConfig)
	return nil, nil
}

// cachingDataSet menghubungkan Storage dengan InMemoryStorage
var cachingDataSet = wire.NewSet(
	service.NewCachingData,
	wire.Bind(new(storage.Storage), new(*storage.CachingData)),
)

// DatabaseStorageSet menghubungkan Storage dengan DatabaseStorage
var databaseStorageSet = wire.NewSet(
	service.NewDatabaseStorage,
	wire.Bind(new(storage.Storage), new(*storage.DatabaseStorage)),
)

// InitializeCachingData menginisialisasi storage menggunakan InMemoryStorageSet
func InitializeCachingData() (storage.Storage, error) {
	wire.Build(cachingDataSet)
	return nil, nil
}

// InitializeDatabaseStorage menginisialisasi storage menggunakan DatabaseStorageSet
func InitializeDatabaseStorage() (storage.Storage, error) {
	wire.Build(databaseStorageSet)
	return nil, nil
}

var notifierSet = wire.NewSet(
	notification.NewEmailService,
	notification.NewSMSService,
	config.NewNotifConfig,
)

// InitializeNotifier menginisialisasi Notifier dengan ketergantungan yang diperlukan
func InitializeNotifier() (*service.Notifier, error) {
	wire.Build(
		notifierSet,
		wire.Struct(new(service.Notifier), "*"),
	)
	return nil, nil
}
