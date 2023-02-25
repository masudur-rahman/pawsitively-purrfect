package services

import (
	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/models/gqtypes"
)

type ShelterService interface {
	ValidateShelter(params gqtypes.ShelterParams) error
	GetShelter(id string) (*models.Shelter, error)
	FindShelters(filter models.Shelter) ([]*models.Shelter, error)
	UserShelters(userID string) ([]*models.Shelter, error)
	CreateShelter(params gqtypes.ShelterParams) (*models.Shelter, error)
	UpdateShelter(shelter gqtypes.ShelterParams) (*models.Shelter, error)
	DeleteShelter(id string) error
}
