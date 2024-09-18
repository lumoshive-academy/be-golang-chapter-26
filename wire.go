//go:build wireinject
// +build wireinject

package main

import (
	"be-golang-chapter-26/greeter"
	"be-golang-chapter-26/service"

	"github.com/google/wire"
)

// func InitializMyService() (*service.Service, error) {
// 	wire.Build(greeter.NewGreeter, service.NewService)
// 	return nil, nil
// }

func InitializMyService(name string) (*service.Service, error) {
	wire.Build(greeter.NewGreeter, service.NewService)
	return nil, nil
}
