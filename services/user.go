package services

import (
	"github.com/masudur-rahman/pawsitively-purrfect/models"
)

type UserService interface {
	ValidateUser(user *models.User) error
	GetUser(id string) (*models.User, error)                           // any logged-in user
	GetUserByName(username string) (*models.User, error)               // any logged-in user
	ListUsers(filter models.User, limit int64) ([]*models.User, error) // mainly for internal uses
	CreateUser(user *models.User) (*models.User, error)                // new user sign up
	UpdateUser(user *models.User) (*models.User, error)                // by logged-in user
	DeleteUser(id string) error                                        // by logged-in user
}
