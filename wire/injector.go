//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/golang-graphql/graph"
	"github.com/golang-graphql/graph/repository"
	"github.com/golang-graphql/graph/service"
	"github.com/golang-graphql/infrastructure"
	"github.com/google/wire"
)

func InitializedServer() *handler.Server {
	wire.Build(
		infrastructure.NewConfig,
		infrastructure.NewMongoDatabase,
		repository.NewUserRepository,
		repository.NewTransactionRepository,
		service.NewUserService,
		service.NewTransactionService,
		graph.NewHandler,
	)
	return nil
}
