package services

import (
	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/models/types"
)

type UserService interface {
	ValidateUser(params types.RegisterParams) error
	GetUser(id string) (*models.User, error)             // any logged-in user
	GetUserByName(username string) (*models.User, error) // any logged-in user
	GetUserByEmail(email string) (*models.User, error)
	ListUsers(filter models.User, limit int64) ([]*models.User, error) // mainly for internal uses
	CreateUser(params types.RegisterParams) (*models.User, error)      // new user sign up
	UpdateUser(user *models.User) (*models.User, error)                // by logged-in user
	DeleteUser(id string) error                                        // by logged-in user
	LoginUser(usernameOrEmail string, passwd string) (*models.User, error)
}
