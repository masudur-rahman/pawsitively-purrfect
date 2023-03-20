package shelter

import (
	"net/http"

	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/models/gqtypes"
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

func (s *shelterService) ValidateShelter(params gqtypes.ShelterParams) error {
	if params.Name == "" {
		return models.StatusError{
			Status:  http.StatusBadRequest,
			Message: "shelter name can't be empty",
		}
	}

	// if params.ID is set, then it's an update operation
	// verify that the shelter owner is the same person
	if params.ID != "" {
		if err := s.ShelterOwnedByUser(params.ID, params.OwnerID); err != nil {
			return err
		}
	}

	shName, err := s.shelterRepo.FindByName(params.Name)
	if err != nil && models.IsErrNotFound(err) {
		return nil
	} else if err != nil {
		return err
	} else if shName.ID != params.ID {
		return models.ErrShelterAlreadyExist{Name: params.Name}
	}

	return nil
}

func (s *shelterService) ShelterOwnedByUser(id, ownerID string) error {
	shelter, err := s.shelterRepo.FindByID(id)
	if err != nil {
		return err
	}

	if shelter.OwnerID != ownerID {
		return models.StatusError{
			Status:  http.StatusBadRequest,
			Message: "shelter isn't owned by logged in user",
		}
	}

	return nil
}

func (s *shelterService) GetShelter(id string) (*models.Shelter, error) {
	return s.shelterRepo.FindByID(id)
}

func (s *shelterService) FindShelters(filter gqtypes.ShelterParams) ([]*models.Shelter, error) {
	sf := models.Shelter{
		ID:                 filter.ID,
		Name:               filter.Name,
		Description:        filter.Description,
		Website:            filter.Website,
		Location:           filter.Location,
		ContactInformation: filter.ContactInformation,
		OwnerID:            filter.OwnerID,
	}
	return s.shelterRepo.FindShelters(sf)
}

func (s *shelterService) UserShelters(userID string) ([]*models.Shelter, error) {
	return s.shelterRepo.FindByOwnerID(userID)
}

func (s *shelterService) CreateShelter(params gqtypes.ShelterParams) (*models.Shelter, error) {
	if err := s.ValidateShelter(params); err != nil {
		return nil, err
	}

	shelter := &models.Shelter{
		Name:               params.Name,
		Description:        params.Description,
		Website:            params.Website,
		Location:           params.Location,
		ContactInformation: params.ContactInformation,
		OwnerID:            params.OwnerID,
	}
	if err := s.shelterRepo.Save(shelter); err != nil {
		return nil, err
	}

	return shelter, nil
}

func (s *shelterService) UpdateShelter(params gqtypes.ShelterParams) (*models.Shelter, error) {
	if params.ID == "" {
		return nil, models.StatusError{
			Status:  http.StatusBadRequest,
			Message: "shelter id missing",
		}
	}

	if err := s.ValidateShelter(params); err != nil {
		return nil, err
	}

	shelter := &models.Shelter{
		ID:                 params.ID,
		Name:               params.Name,
		Description:        params.Description,
		Website:            params.Website,
		Location:           params.Location,
		ContactInformation: params.ContactInformation,
		OwnerID:            params.OwnerID,
	}
	if err := s.shelterRepo.Update(shelter); err != nil {
		return nil, err
	}
	return shelter, nil
}

func (s *shelterService) IncreasePetCount(shelterID string) error {
	shelter, err := s.shelterRepo.FindByID(shelterID)
	if err != nil {
		return err
	}

	shelter.NumberOfPets++
	return s.shelterRepo.Update(shelter)
}

func (s *shelterService) DecreasePetCount(shelterID string) error {
	shelter, err := s.shelterRepo.FindByID(shelterID)
	if err != nil {
		return err
	}

	shelter.NumberOfPets--
	return s.shelterRepo.Update(shelter)
}

func (s *shelterService) DeleteShelter(id string) error {
	return s.shelterRepo.Delete(id)
}
