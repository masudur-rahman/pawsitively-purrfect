package http

import (
	"github.com/masudur-rahman/pawsitively-purrfect/api/http/handlers"
	"github.com/masudur-rahman/pawsitively-purrfect/api/http/middlewares"
	"github.com/masudur-rahman/pawsitively-purrfect/services/all"

	"github.com/flamego/binding"
	"github.com/flamego/flamego"
	"golang.org/x/time/rate"
)

func Routes(svc *all.Services) *flamego.Flame {
	f := flamego.Classic()
	f.Get("/", func() string {
		return "Hello, Flamego!\n"
	})

	f.Use(middlewares.RateLimiter(rate.NewLimiter(10, 20)))

	f.Map(svc)
	f.Use(middlewares.Sessioner())
	f.Use(middlewares.CSRFer())
	f.Use(middlewares.ReqPurrfectContext())

	f.Any("/graphql", binding.JSON(handlers.RequestOptions{}), handlers.ServeGraphQL)

	return f
}
