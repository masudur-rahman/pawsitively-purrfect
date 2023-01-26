package pet

import (
	"fmt"

	"github.com/masudur-rahman/pawsitively-purrfect/models"
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
		return fmt.Errorf("could not find pet with ID %s: %w", petID, err)
	}

	if pet.AdoptionStatus == models.PetAdopted {
		return fmt.Errorf("pet already adopted")
	}

	_, err = p.userRepo.FindByID(userID)
	if err != nil {
		return fmt.Errorf("pet with ID %s already adopted", petID)
	}

	pet.AdoptionStatus = models.PetAdopted
	pet.CurrentOwnerID = userID
	pet.OriginShelterID, pet.ShelterID = pet.ShelterID, ""
	err = p.petRepo.Update(pet)
	if err != nil {
		return fmt.Errorf("error while updating the pet: %v", err)
	}
	return nil
}

func (p *petService) AddPetToShelter(shelterID string, pet *models.Pet) error {
	if shelterID == "" || pet == nil {
		return fmt.Errorf("invalid input")
	}

	// TODO: might need to check if the shelterID is valid
	pet.ShelterID = shelterID

	err := p.petRepo.Save(pet)
	if err != nil {
		return fmt.Errorf("failed to add pet to shelter: %w", err)
	}

	return nil
}

func (p *petService) GetPetByID(id string) (*models.Pet, error) {
	return p.petRepo.FindByID(id)
}

func (p *petService) UpdatePet(pet *models.Pet) error {
	return p.petRepo.Update(pet)
}

func (p *petService) DeletePet(id string) error {
	return p.petRepo.Delete(id)
}
