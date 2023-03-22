package resolvers

import (
	"errors"
	"net/http"

	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/models/gqtypes"
	"github.com/masudur-rahman/pawsitively-purrfect/pkg"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) GetPet(p graphql.ResolveParams) (interface{}, error) {
	if !r.ctx.IsAuthenticated() {
		return nil, models.ErrUserNotAuthenticated{}
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
	if !r.ctx.IsAuthenticated() {
		return nil, models.ErrUserNotAuthenticated{}
	}

	shelterID, ok := p.Args["shelterID"].(string)
	if !ok {
		return nil, errors.New("invalid argument")
	}

	_, err := r.svc.Shelter.GetShelter(shelterID)
	if err != nil {
		return nil, err
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
	if !r.ctx.IsAuthenticated() {
		return nil, models.ErrUserNotAuthenticated{}
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

	if err = r.svc.Shelter.IncreasePetCount(params.ShelterID); err != nil {
		return nil, err
	}

	return pet.APIFormat(), nil
}

func (r *Resolver) UpdatePet(p graphql.ResolveParams) (interface{}, error) {
	if !r.ctx.IsAuthenticated() {
		return nil, models.ErrUserNotAuthenticated{}
	}

	params := gqtypes.PetParams{}
	if err := pkg.ParseInto(p.Args, &params); err != nil {
		return nil, err
	}

	pet, err := r.svc.Pet.GetPetByID(params.ID)
	if err != nil {
		return nil, err
	}

	owner := pet.CurrentOwnerID
	if pet.ShelterID != "" {
		shelter, err := r.svc.Shelter.GetShelter(pet.ShelterID)
		if err != nil {
			return nil, err
		}
		owner = shelter.OwnerID
	}

	if owner != r.ctx.GetLoggedInUserID() {
		return nil, models.StatusError{
			Status:  http.StatusBadRequest,
			Message: "pet doesn't belong to current user",
		}
	}

	pet, err = r.svc.Pet.UpdatePet(params)
	if err != nil {
		return nil, err
	}

	return pet.APIFormat(), nil
}

func (r *Resolver) AdoptPet(p graphql.ResolveParams) (interface{}, error) {
	if !r.ctx.IsAuthenticated() {
		return nil, models.ErrUserNotAuthenticated{}
	}

	id, ok := p.Args["id"].(string)
	if !ok {
		return nil, errors.New("invalid argument")
	}
	pet, err := r.svc.Pet.GetPetByID(id)
	if err != nil {
		return nil, err
	}

	if pet.ShelterID == "" {
		return nil, models.StatusError{
			Status:  http.StatusBadRequest,
			Message: "no shelter owns this pet",
		}
	}

	if err = r.svc.Pet.AdoptPet(r.ctx.GetLoggedInUserID(), pet.ID); err != nil {
		return nil, err
	}

	if err = r.svc.Shelter.DecreasePetCount(pet.ShelterID); err != nil {
		return nil, err
	}

	return nil, nil
}
