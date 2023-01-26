package services

import "github.com/masudur-rahman/pawsitively-purrfect/models"

type PetService interface {
	AdoptPet(userID string, petID string) error
	AddPetToShelter(shelterID string, pet *models.Pet) error
	GetPetByID(id string) (*models.Pet, error)
	UpdatePet(pet *models.Pet) error
	DeletePet(id string) error
}
