package services

import "github.com/masudur-rahman/pawsitively-purrfect/models"

type ShelterService interface {
	GetShelter(id string) (*models.Shelter, error)
	FindShelters(filter models.Shelter) ([]*models.Shelter, error)
	UserShelters(userID string) ([]*models.Shelter, error)
	CreateShelter(shelter *models.Shelter) (*models.Shelter, error)
	UpdateShelter(shelter *models.Shelter) (*models.Shelter, error)
	DeleteShelter(id string) error
}
