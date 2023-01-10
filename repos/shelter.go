package repos

import "github.com/masudur-rahman/pawsitively-purrfect/models"

type ShelterRepository interface {
	FindByID(id string) (*models.Shelter, error)
	FindByName(name string) (*models.Shelter, error)
	FindByLocation(location string) ([]*models.Shelter, error)
	FindByOwnerID(id int64) ([]*models.Shelter, error)
	FindShelters(filter models.Shelter) ([]*models.Shelter, error)
	Save(shelter *models.Shelter) error
	Update(shelter *models.Shelter) error
	Delete(id string) error
}
