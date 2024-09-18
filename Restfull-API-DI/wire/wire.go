//go:build wireinject
// +build wireinject

package wire

import (
	"be-golang-chapter-26/Restfull-API-DI/database"
	"be-golang-chapter-26/Restfull-API-DI/handler"
	"be-golang-chapter-26/Restfull-API-DI/repository"
	"be-golang-chapter-26/Restfull-API-DI/service"

	"github.com/google/wire"
)

var productHandlerSet = wire.NewSet(
	database.NewPostgresDB,
	repository.NewProductRepository,
	service.NewProductService,
	handler.NewProductHandler,
)

func InitializProductHandler() *handler.ProductHandler {
	wire.Build(productHandlerSet)
	return nil
}
