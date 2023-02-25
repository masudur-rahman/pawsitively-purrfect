package resolvers

import (
	"errors"

	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/models/gqtypes"
	"github.com/masudur-rahman/pawsitively-purrfect/pkg"

	"github.com/graphql-go/graphql"
)

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

func (r *Resolver) AddShelter(p graphql.ResolveParams) (interface{}, error) {
	if !r.IsAuthenticated() {
		return nil, models.ErrUserNotAuthenticated
	}

	params := gqtypes.ShelterParams{}
	if err := pkg.ParseInto(p.Args, &params); err != nil {
		return nil, err
	}

	params.OwnerID = r.ctx.User.ID
	if err := r.svc.Shelter.ValidateShelter(params); err != nil {
		return nil, err
	}

	shelter, err := r.svc.Shelter.CreateShelter(params)
	if err != nil {
		return nil, err
	}

	return shelter.APIFormat(), nil
}
