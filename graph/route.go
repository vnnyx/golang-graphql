package graph

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/golang-graphql/graph/generated"
	"github.com/golang-graphql/graph/service"
)

type Route struct {
	Router             *chi.Mux
	UserService        service.UserService
	TransactionService service.TransactionService
}

func New(router *chi.Mux, userService service.UserService, transactionService service.TransactionService) *Route {
	return &Route{
		Router:             router,
		UserService:        userService,
		TransactionService: transactionService,
	}
}

func (r *Route) InitRoute() {
	r.Router.Post("/query", graphQLHandler(r.UserService, r.TransactionService))
	r.Router.Get("/", playgroundQLHandler("/query"))
}

func graphQLHandler(userService service.UserService, transactionService service.TransactionService) http.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{
		UserService:        userService,
		TransactionService: transactionService,
	}}))

	return h.ServeHTTP
}

func playgroundQLHandler(endpoint string) http.HandlerFunc {
	playgroundHandler := playground.Handler("GraphQL playground", endpoint)

	return playgroundHandler
}
