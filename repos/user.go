package repos

import "github.com/masudur-rahman/pawsitively-purrfect/models"

type UserRepository interface {
	FindByID(id string) (*models.User, error)
	FindByName(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindUsers(filter models.User) ([]*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id string) error
}
