package services

import (
	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/models/gqtypes"
)

type PetService interface {
	AdoptPet(userID string, petID string) error
	AddPetToShelter(params gqtypes.PetParams) (*models.Pet, error)
	GetPetByID(id string) (*models.Pet, error)
	FindPets(params gqtypes.PetParams) ([]*models.Pet, error)
	ListShelterPets(shelterID string) ([]*models.Pet, error)
	ListPetsOwnedByUser(userID string) ([]*models.Pet, error)
	GetPetOwnerID(petID string) (string, error)
	UpdatePet(params gqtypes.PetParams) (*models.Pet, error)
	DeletePet(id string) error
}
