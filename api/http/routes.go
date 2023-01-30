package http

import (
	"github.com/masudur-rahman/pawsitively-purrfect/api/graphql/resolvers"
	"github.com/masudur-rahman/pawsitively-purrfect/api/http/handlers"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/flamego/binding"
	"github.com/flamego/flamego"
)

func Routes(resolver *resolvers.Resolver) *flamego.Flame {
	f := flamego.Classic()
	f.Get("/", func() string {
		return "Hello, Flamego!\n"
	})

	f.Map(resolver)
	f.Any("/graphql", binding.JSON(handlers.RequestOptions{}), handlers.ServeGraphQL)
	f.Any("/play", playground.Handler("GraphQL playground", "/graphql"))

	return f
}
