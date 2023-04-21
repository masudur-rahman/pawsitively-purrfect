package resolvers

import (
	"errors"
	"net/http"

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

	return shelter.APIFormat(), nil
}

func (r *Resolver) ListShelters(p graphql.ResolveParams) (interface{}, error) {
	if !r.ctx.IsAuthenticated() {
		return nil, models.ErrUserNotAuthenticated{}
	}

	params := gqtypes.ShelterParams{}
	if err := pkg.ParseInto(p.Args, &params); err != nil {
		return nil, err
	}

	shelters, err := r.svc.Shelter.FindShelters(params)
	if err != nil {
		return nil, err
	}

	apiShelters := make([]gqtypes.Shelter, 0, len(shelters))
	for _, shelter := range shelters {
		apiShelters = append(apiShelters, shelter.APIFormat())
	}

	return apiShelters, nil
}

func (r *Resolver) AddShelter(p graphql.ResolveParams) (interface{}, error) {
	if !r.ctx.IsAuthenticated() {
		return nil, models.ErrUserNotAuthenticated{}
	}

	params := gqtypes.ShelterParams{}
	if err := pkg.ParseInto(p.Args, &params); err != nil {
		return nil, err
	}

	params.OwnerID = r.ctx.GetLoggedInUserID()
	if err := r.svc.Shelter.ValidateShelter(params); err != nil {
		return nil, err
	}

	shelter, err := r.svc.Shelter.CreateShelter(params)
	if err != nil {
		return nil, err
	}

	return shelter.APIFormat(), nil
}

func (r *Resolver) UpdateShelter(p graphql.ResolveParams) (interface{}, error) {
	if !r.ctx.IsAuthenticated() {
		return nil, models.ErrUserNotAuthenticated{}
	}

	params := gqtypes.ShelterParams{}
	if err := pkg.ParseInto(p.Args, &params); err != nil {
		return nil, err
	}

	shelter, err := r.svc.Shelter.GetShelter(params.ID)
	if err != nil {
		return nil, err
	}
	if shelter.OwnerID != r.ctx.GetLoggedInUserID() {
		return nil, models.StatusError{
			Status:  http.StatusBadRequest,
			Message: "user don't own the shelter",
		}
	}

	shelter, err = r.svc.Shelter.UpdateShelter(params)
	if err != nil {
		return nil, err
	}

	return shelter.APIFormat(), nil
}
