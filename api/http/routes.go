package http

import (
	"net/http"

	"github.com/masudur-rahman/pawsitively-purrfect/api/graphql/resolvers"
	"github.com/masudur-rahman/pawsitively-purrfect/api/graphql/schema"
	"github.com/masudur-rahman/pawsitively-purrfect/api/http/handlers"
	"github.com/masudur-rahman/pawsitively-purrfect/api/http/middlewares"
	"github.com/masudur-rahman/pawsitively-purrfect/services/all"

	"github.com/flamego/binding"
	"github.com/flamego/flamego"
	"github.com/graphql-go/handler"
)

func Routes(svc *all.Services) *flamego.Flame {
	f := flamego.Classic()
	f.Get("/", func() string {
		return "Hello, Flamego!\n"
	})

	f.Map(svc)
	f.Use(middlewares.Sessioner())
	f.Use(middlewares.CSRFer())
	f.Use(middlewares.ReqPurrfectContext())

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
