package resolvers

import (
	"errors"
	"net/http"

	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/models/gqtypes"
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
	return user.APIFormat(), nil
}

func (r *Resolver) RegisterUser(p graphql.ResolveParams) (interface{}, error) {
	params := gqtypes.RegisterParams{}
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
	params := gqtypes.LoginParams{}
	err := pkg.ParseInto(p.Args, &params)
	if err != nil {
		return nil, err
	}

	user, err := r.svc.User.LoginUser(params.Username, params.Password)
	if err != nil {
		return nil, err
	}

	return user.APIFormat(), err
}

func (r *Resolver) Profile(p graphql.ResolveParams) (interface{}, error) {
	if !r.ctx.IsAuthenticated() {
		return nil, models.ErrUserNotAuthenticated{}
	}

	return r.ctx.User.APIFormat(), nil
}

func (r *Resolver) UpdateProfile(p graphql.ResolveParams) (interface{}, error) {
	if !r.ctx.IsAuthenticated() {
		return nil, models.ErrUserNotAuthenticated{}
	}

	params := gqtypes.UserParams{}
	if err := pkg.ParseInto(p.Args, &params); err != nil {
		return nil, err
	}

	// TODO: need to add username update support
	if params.ID != r.ctx.GetLoggedInUserID() || params.Username != r.ctx.User.Username {
		return nil, models.StatusError{
			Status:  http.StatusBadRequest,
			Message: "can't update id or username",
		}
	}

	user, err := r.svc.User.UpdateUser(params)
	if err != nil {
		return nil, err
	}

	return user.APIFormat(), nil
}
