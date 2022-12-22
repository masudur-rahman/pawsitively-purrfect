package repos

import "github.com/masudur-rahman/pawsitively-purrfect/models"

type UserRepository interface {
	FindByID(id string) (*models.User, error)
	FindByNme(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id string) error
}
