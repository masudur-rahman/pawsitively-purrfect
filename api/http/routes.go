package http

import (
	"github.com/masudur-rahman/pawsitively-purrfect/api/http/handlers"
	"github.com/masudur-rahman/pawsitively-purrfect/api/http/middlewares"
	"github.com/masudur-rahman/pawsitively-purrfect/services/all"
	"github.com/masudur-rahman/pawsitively-purrfect/templates"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/flamego/binding"
	"github.com/flamego/flamego"
	"github.com/flamego/template"
	"golang.org/x/time/rate"
)

func Routes(svc *all.Services) *flamego.Flame {
	f := flamego.Classic()
	f.Get("/", func() string {
		return "Hello, Flamego!\n"
	})

	f.Use(middlewares.RateLimiter(rate.NewLimiter(10, 20)))

	f.Map(svc)

	fs, err := template.EmbedFS(templates.Templates, ".", []string{".tmpl"})
	if err != nil {
		panic(err)
	}
	f.Use(template.Templater(
		template.Options{
			FileSystem: fs,
		},
	))

	f.Use(middlewares.Sessioner())
	f.Use(middlewares.CSRFer())
	f.Use(middlewares.ReqPurrfectContext())

	f.Combo("/user/login").
		Get(handlers.Login).
		Post(handlers.LoginPost)
	f.Combo("/user/register").
		Get(handlers.Register).
		Post(handlers.RegisterPost)

	f.Any("/graphql", binding.JSON(handlers.RequestOptions{}), handlers.ServeGraphQL)
	f.Any("/playground", playground.Handler("Pawsitively Purrfect", "/graphql"))

	return f
}
