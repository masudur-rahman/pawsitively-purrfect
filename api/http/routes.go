package http

import (
	"github.com/masudur-rahman/pawsitively-purrfect/api/graphql/resolvers"
	"github.com/masudur-rahman/pawsitively-purrfect/api/http/handlers"

	"github.com/flamego/binding"
	"github.com/flamego/flamego"
)

func Routes(resolver *resolvers.Resolver) *flamego.Flame {
	f := flamego.Classic()
	f.Get("/", func() string {
		return "Hello, Flamego!"
	})

	f.Map(resolver)
	f.Any("/graphql", binding.JSON(handlers.RequestOptions{}), handlers.ServeGraphQL)

	return f
}
