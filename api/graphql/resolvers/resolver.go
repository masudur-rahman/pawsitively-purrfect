package resolvers

import (
	"github.com/masudur-rahman/pawsitively-purrfect/api/http/middlewares"
	"github.com/masudur-rahman/pawsitively-purrfect/services/all"
)

type Resolver struct {
	ctx *middlewares.PurrfectContext
	svc *all.Services
}

func NewResolver(ctx *middlewares.PurrfectContext, svc *all.Services) *Resolver {
	return &Resolver{ctx: ctx, svc: svc}
}
