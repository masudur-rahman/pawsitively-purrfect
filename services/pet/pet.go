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
	petRepo         repos.PetRepository
	userRepo        repos.UserRepository
	petAdoptionRepo repos.PetAdoptionRepository
}

var _ services.PetService = &petService{}

func NewPetService(petRepo repos.PetRepository, userRepo repos.UserRepository, paRepo repos.PetAdoptionRepository) *petService {
	return &petService{petRepo: petRepo, userRepo: userRepo, petAdoptionRepo: paRepo}
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
	if err = p.petAdoptionRepo.AddPetAdoption(petID, userID); err != nil {
		return err
	}

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
		Type:           models.PetType(params.Type),
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

func (p *petService) ListPetsOwnedByUser(userID string) ([]*models.Pet, error) {
	ids, err := p.petAdoptionRepo.ListPetsAdoptedByUser(userID)
	if err != nil {
		return nil, err
	}
	return p.petRepo.FindByIDs(ids)
}

func (p *petService) ListShelterPets(shelterID string) ([]*models.Pet, error) {
	return p.petRepo.FindByShelterID(shelterID)
}

func (p *petService) GetPetOwnerID(petID string) (string, error) {
	pet, err := p.petRepo.FindByID(petID)
	if err != nil {
		return "", err
	}
	if pet.AdoptionStatus != models.PetAdopted {
		return "", models.StatusError{
			Status:  http.StatusBadRequest,
			Message: "pet doesn't have an owner",
		}
	}

	return p.petAdoptionRepo.GetPetOwner(petID)
}

func (p *petService) UpdatePet(params gqtypes.PetParams) (*models.Pet, error) {
	pet, err := p.petRepo.FindByID(params.ID)
	if err != nil {
		return nil, err
	}

	pet.Name = params.Name
	pet.Breed = params.Breed
	pet.Gender = params.Gender

	if err = p.petRepo.Update(pet); err != nil {
		return nil, err
	}

	return pet, nil
}

func (p *petService) DeletePet(id string) error {
	return p.petRepo.Delete(id)
}
