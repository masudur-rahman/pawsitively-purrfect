package pet

import (
	"fmt"
	"net/http"

	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/models/gqtypes"
	"github.com/masudur-rahman/pawsitively-purrfect/repos"
	"github.com/masudur-rahman/pawsitively-purrfect/services"
)

type petService struct {
	petRepo  repos.PetRepository
	userRepo repos.UserRepository
}

var _ services.PetService = &petService{}

func NewPetService(petRepo repos.PetRepository, userRepo repos.UserRepository) *petService {
	return &petService{petRepo: petRepo, userRepo: userRepo}
}

func (p *petService) AdoptPet(userID string, petID string) error {
	pet, err := p.petRepo.FindByID(petID)
	if err != nil {
		return err
	}

	if pet.AdoptionStatus == models.PetAdopted {
		return models.StatusError{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("pet with ID %s already adopted", petID),
		}
	}

	_, err = p.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	pet.AdoptionStatus = models.PetAdopted
	pet.CurrentOwnerID = userID
	pet.OriginShelterID, pet.ShelterID = pet.ShelterID, ""
	err = p.petRepo.Update(pet)
	if err != nil {
		return models.StatusError{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("error while updating the pet: %v", err),
		}
	}
	return nil
}

func (p *petService) AddPetToShelter(params gqtypes.PetParams) (*models.Pet, error) {
	if params.ShelterID == "" {
		return nil, models.StatusError{
			Status:  http.StatusBadRequest,
			Message: "shelterID must be non empty",
		}
	}

	pet := &models.Pet{
		Name:           params.Name,
		Breed:          params.Breed,
		Gender:         params.Gender,
		AdoptionStatus: models.PetAvailable,
		ShelterID:      params.ShelterID,
	}

	err := p.petRepo.Save(pet)
	if err != nil {
		return nil, models.StatusError{Message: fmt.Sprintf("failed to add pet to shelter: %v", err)}
	}

	return pet, nil
}

func (p *petService) GetPetByID(id string) (*models.Pet, error) {
	return p.petRepo.FindByID(id)
}

func (p *petService) ListShelterPets(shelterID string) ([]*models.Pet, error) {
	return p.petRepo.FindByShelterID(shelterID)
}

func (p *petService) UpdatePet(pet *models.Pet) error {
	return p.petRepo.Update(pet)
}

func (p *petService) DeletePet(id string) error {
	return p.petRepo.Delete(id)
}
