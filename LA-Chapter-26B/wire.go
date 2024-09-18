//go:build wireinject
// +build wireinject

package main

import (
	"be-golang-chapter-26/LA-Chapter-26B/greeter"
	"be-golang-chapter-26/LA-Chapter-26B/service"

	"github.com/google/wire"
)

func InitializMyService() *service.Service {
	wire.Build(greeter.NewGreeter, service.NewService)
	return nil
}
