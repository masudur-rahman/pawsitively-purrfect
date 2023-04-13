package http

import (
	"github.com/masudur-rahman/pawsitively-purrfect/api/http/handlers"
	"github.com/masudur-rahman/pawsitively-purrfect/api/http/middlewares"
	"github.com/masudur-rahman/pawsitively-purrfect/models/gqtypes"
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

	f.Get("/", handlers.Home)

	f.Group("/user", func() {
		f.Combo("/login").Get(handlers.Login).
			Post(binding.Form(gqtypes.LoginParams{}), handlers.LoginPost)
		f.Combo("/register").Get(handlers.Register).
			Post(binding.Form(gqtypes.RegisterParams{}), handlers.RegisterPost)
	}, middlewares.ReqSignedOut())

	f.Get("/logout", middlewares.ReqAuth(), handlers.Logout)
	f.Get("/{name}", middlewares.ReqAuth(), handlers.Profile)

	f.Any("/graphql", binding.JSON(handlers.RequestOptions{}), handlers.ServeGraphQL)
	f.Any("/playground", playground.Handler("Pawsitively Purrfect", "/graphql"))

	return f
}
