package resolvers

import (
	"errors"

	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/services"

	"github.com/graphql-go/graphql"
)

type Resolver struct {
	us services.UserService
	ss services.ShelterService
	ps services.PetService
}

func NewResolver(us services.UserService, ss services.ShelterService, ps services.PetService) *Resolver {
	return &Resolver{us: us, ss: ss, ps: ps}
}

func (r *Resolver) GetUser(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var user *models.User
	if id, ok := p.Args["id"].(string); ok {
		user, err = r.us.GetUser(id)
	} else if name, ok := p.Args["name"].(string); ok {
		user, err = r.us.GetUserByName(name)
	} else {
		return nil, errors.New("invalid argument")
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Resolver) GetShelter(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(string)
	if !ok {
		return nil, errors.New("invalid argument")
	}
	shelter, err := r.ss.GetShelter(id)
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
	pet, err := r.ps.GetPetByID(id)
	if err != nil {
		return nil, err
	}

	return pet, nil
}
