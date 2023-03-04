package resolvers

import (
	"errors"

	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/models/gqtypes"
	"github.com/masudur-rahman/pawsitively-purrfect/pkg"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) GetPet(p graphql.ResolveParams) (interface{}, error) {
	if !r.IsAuthenticated() {
		return nil, models.ErrUserNotAuthenticated
	}

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

func (r *Resolver) ListPets(p graphql.ResolveParams) (interface{}, error) {
	if !r.IsAuthenticated() {
		return nil, models.ErrUserNotAuthenticated
	}

	shelterID, ok := p.Args["shelterID"].(string)
	if !ok {
		return nil, errors.New("invalid argument")
	}

	pets, err := r.svc.Pet.ListShelterPets(shelterID)
	if err != nil {
		return nil, err
	}

	apiPets := make([]gqtypes.Pet, 0, len(pets))
	for _, pet := range pets {
		apiPets = append(apiPets, pet.APIFormat())
	}

	return apiPets, nil
}

func (r *Resolver) AddPetNewPet(p graphql.ResolveParams) (interface{}, error) {
	if !r.IsAuthenticated() {
		return nil, models.ErrUserNotAuthenticated
	}

	params := gqtypes.PetParams{}
	if err := pkg.ParseInto(p.Args, &params); err != nil {
		return nil, err
	}

	if err := r.svc.Shelter.ShelterOwnedByUser(params.ShelterID, r.ctx.GetLoggedInUserID()); err != nil {
		return nil, err
	}

	pet, err := r.svc.Pet.AddPetToShelter(params)
	if err != nil {
		return nil, err
	}

	return pet.APIFormat(), nil
}
