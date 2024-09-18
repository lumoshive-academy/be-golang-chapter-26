// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"be-golang-chapter-26/greeter"
	"be-golang-chapter-26/service"
)

// Injectors from wire.go:

func InitializMyService(name string) (*service.Service, error) {
	greeterGreeter := greeter.NewGreeter(name)
	serviceService, err := service.NewService(greeterGreeter)
	if err != nil {
		return nil, err
	}
	return serviceService, nil
}
