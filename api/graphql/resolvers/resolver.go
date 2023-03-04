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

func (r *Resolver) IsAuthenticated() bool {
	if r.ctx.IsSigned && (r.ctx.IsBasicAuth || r.ctx.IsValidCSRF) {
		return true
	}

	return false
}
