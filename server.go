package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/golang-graphql/graph"
	"github.com/golang-graphql/graph/generated"
	"github.com/golang-graphql/graph/repository"
	"github.com/golang-graphql/graph/service"
	"github.com/golang-graphql/infrastructure"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	config, _ := infrastructure.NewConfig()
	mongo := infrastructure.NewMongoDatabase(*config)

	userRepository := repository.NewUserRepository(mongo)
	transactionRepository := repository.NewTransactionRepository(mongo)

	userService := service.NewUserService(&userRepository)
	transactionService := service.NewTransactionService(&transactionRepository)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		UserService:        userService,
		TransactionService: transactionService,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
