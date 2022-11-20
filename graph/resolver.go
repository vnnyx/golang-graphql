package graph

import "github.com/golang-graphql/graph/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	service.UserService
	service.TransactionService
}
