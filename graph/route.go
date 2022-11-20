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
	Router  *chi.Mux
	Service service.UserService
}

func New(router *chi.Mux, service service.UserService) *Route {
	return &Route{
		Router:  router,
		Service: service,
	}
}

func (r *Route) InitRoute() {
	r.Router.Post("/query", graphQLHandler(r.Service))
	// r.Router.Get("/", playgroundQLHandler("/query"))
}

func graphQLHandler(service service.UserService) http.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{
		UserService: service,
	}}))

	return h.ServeHTTP
}

func playgroundQLHandler(endpoint string) http.HandlerFunc {
	playgroundHandler := playground.Handler("GraphQL playground", endpoint)

	return playgroundHandler
}
