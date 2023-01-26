package shelter

import (
	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/repos"
	"github.com/masudur-rahman/pawsitively-purrfect/services"
)

type shelterService struct {
	shelterRepo repos.ShelterRepository
}

var _ services.ShelterService = &shelterService{}

func NewShelterService(shelterRepo repos.ShelterRepository) *shelterService {
	return &shelterService{
		shelterRepo: shelterRepo,
	}
}

func (s *shelterService) GetShelter(id string) (*models.Shelter, error) {
	return s.shelterRepo.FindByID(id)
}

func (s *shelterService) FindShelters(filter models.Shelter) ([]*models.Shelter, error) {
	return s.shelterRepo.FindShelters(filter)
}

func (s *shelterService) UserShelters(userID string) ([]*models.Shelter, error) {
	return s.shelterRepo.FindByOwnerID(userID)
}

func (s *shelterService) CreateShelter(shelter *models.Shelter) (*models.Shelter, error) {
	if err := s.shelterRepo.Save(shelter); err != nil {
		return nil, err
	}

	return shelter, nil
}

func (s *shelterService) UpdateShelter(shelter *models.Shelter) (*models.Shelter, error) {
	if err := s.shelterRepo.Update(shelter); err != nil {
		return nil, err
	}
	return shelter, nil
}

func (s *shelterService) DeleteShelter(id string) error {
	return s.shelterRepo.Delete(id)
}
