package resolvers

import (
	"errors"

	"github.com/masudur-rahman/pawsitively-purrfect/models"

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
	user := new(models.User)
	if username, ok := p.Args["username"].(string); ok {
		user.Username = username
	}
	if email, ok := p.Args["email"].(string); ok {
		user.Email = email
	}
	if password, ok := p.Args["password"].(string); ok {
		user.PasswordHash = password
	}

	user, err := r.us.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
