package repos

import "github.com/masudur-rahman/pawsitively-purrfect/models"

type PetRepository interface {
	FindByID(id string) (*models.Pet, error)
	FindByIDs(ids []string) ([]*models.Pet, error)
	FindByType(typ models.PetType) ([]*models.Pet, error)
	FindByBreed(breed string) ([]*models.Pet, error)
	FindByGender(gender string) ([]*models.Pet, error)
	FindByAdoptionStatus(status models.PetAdoptionStatus) ([]*models.Pet, error)
	FindByShelterID(id string) ([]*models.Pet, error)
	FindPets(filter models.Pet) ([]*models.Pet, error)
	Save(pet *models.Pet) error
	Update(pet *models.Pet) error
	Delete(id string) error
}
