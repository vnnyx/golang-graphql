package graph

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/golang-graphql/graph/generated"
	"github.com/golang-graphql/graph/service"
)

func NewHandler(userService service.UserService, transactionService service.TransactionService) *handler.Server {
	return handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{
		UserService:        userService,
		TransactionService: transactionService,
	}}))
}
