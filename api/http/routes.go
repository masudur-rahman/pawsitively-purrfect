package http

import (
	"net/http"

	"github.com/masudur-rahman/pawsitively-purrfect/api/graphql/resolvers"
	"github.com/masudur-rahman/pawsitively-purrfect/api/graphql/schema"
	"github.com/masudur-rahman/pawsitively-purrfect/api/http/handlers"

	"github.com/flamego/binding"
	"github.com/flamego/flamego"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func Routes(schemas graphql.Schema) *flamego.Flame {
	f := flamego.Classic()
	f.Get("/", func() string {
		return "Hello, Flamego!\n"
	})

	f.Map(schemas)

	f.Any("/graphql", binding.JSON(handlers.RequestOptions{}), handlers.ServeGraphQL)

	return f
}

func Graph(resolver *resolvers.Resolver) {
	schemas, err := schema.PurrfectSchema(resolver)
	if err != nil {
		panic(err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schemas,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
