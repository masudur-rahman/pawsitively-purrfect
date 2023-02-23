package resolvers

import (
	"errors"

	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/models/types"
	"github.com/masudur-rahman/pawsitively-purrfect/pkg"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) GetUser(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var user *models.User
	if id, ok := p.Args["id"].(string); ok {
		user, err = r.svc.User.GetUser(id)
	} else if name, ok := p.Args["name"].(string); ok {
		user, err = r.svc.User.GetUserByName(name)
	} else {
		return nil, errors.New("invalid argument")
	}
	if err != nil {
		return nil, err
	}
	return user.APIUser(), nil
}

func (r *Resolver) RegisterUser(p graphql.ResolveParams) (interface{}, error) {
	params := types.RegisterParams{}
	if err := pkg.ParseInto(p.Args, &params); err != nil {
		return nil, err
	}

	user, err := r.svc.User.CreateUser(params)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Resolver) Login(p graphql.ResolveParams) (interface{}, error) {
	params := types.LoginParams{}
	err := pkg.ParseInto(p.Args, &params)
	if err != nil {
		return nil, err
	}

	user, err := r.svc.User.LoginUser(params.Username, params.Password)
	if err != nil {
		return nil, err
	}

	return user.APIUser(), err
}
