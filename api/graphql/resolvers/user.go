package resolvers

import (
	"errors"
	"strings"

	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/models/types"
	"github.com/masudur-rahman/pawsitively-purrfect/pkg"

	"github.com/graphql-go/graphql"
)

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

func (r *Resolver) RegisterUser(p graphql.ResolveParams) (interface{}, error) {
	params := types.RegisterParams{}
	if err := pkg.ParseInto(p.Args, &params); err != nil {
		return nil, err
	}

	user, err := r.us.CreateUser(params)
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

	var user *models.User
	if strings.Contains(params.Username, "@") {
		user, err = r.us.GetUserByEmail(params.Username)
		if err != nil {
			return nil, err
		}
	} else {
		user, err = r.us.GetUserByName(params.Username)
		if err != nil {
			return nil, err
		}
	}

	if !pkg.CheckPasswordHash(params.Password, user.PasswordHash) {
		return nil, errors.New("username or password is invalid")
	}

	//TODO: Token generation, set to cookie

	return nil, err
}
