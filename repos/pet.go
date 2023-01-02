package repos

import "github.com/masudur-rahman/pawsitively-purrfect/models"

type PetRepository interface {
	FindByID(id string) (*models.Pet, error)
	FindByBreed(breed string) ([]*models.Pet, error)
	FindByGender(gender string) ([]*models.Pet, error)
	FindByAdoptionStatus(status string) ([]*models.Pet, error)
	FindByShelterID(id string) ([]*models.Pet, error)
	FindByCurrentOwnerID(id string) ([]*models.Pet, error)
	FindByOriginShelterID(id string) ([]*models.Pet, error)
	Save(pet *models.Pet) error
	Update(pet *models.Pet) error
	Delete(id string) error
}
