package resolvers

import (
	"errors"

	"github.com/masudur-rahman/pawsitively-purrfect/api/http/middlewares"
	"github.com/masudur-rahman/pawsitively-purrfect/services/all"

	"github.com/graphql-go/graphql"
)

type Resolver struct {
	ctx *middlewares.PurrfectContext
	svc *all.Services
}

func NewResolver(ctx *middlewares.PurrfectContext, svc *all.Services) *Resolver {
	return &Resolver{ctx: ctx, svc: svc}
}

func (r *Resolver) GetShelter(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(string)
	if !ok {
		return nil, errors.New("invalid argument")
	}
	shelter, err := r.svc.Shelter.GetShelter(id)
	if err != nil {
		return nil, err
	}

	return shelter, nil
}

func (r *Resolver) GetPet(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(string)
	if !ok {
		return nil, errors.New("invalid argument")
	}
	pet, err := r.svc.Pet.GetPetByID(id)
	if err != nil {
		return nil, err
	}

	return pet, nil
}
