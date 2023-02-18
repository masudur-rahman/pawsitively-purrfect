package shelter

import (
	"errors"
	"fmt"

	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/nosql"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/logr"
	"github.com/masudur-rahman/pawsitively-purrfect/models"
)

type NoSQLShelterRepository struct {
	db     nosql.Database
	logger logr.Logger
}

func NewNoSQLShelterRepository(db nosql.Database, logger logr.Logger) *NoSQLShelterRepository {
	return &NoSQLShelterRepository{
		db:     db.Collection("shelter"),
		logger: logger,
	}
}

func (s *NoSQLShelterRepository) FindByID(id string) (*models.Shelter, error) {
	var shelter models.Shelter
	found, err := s.db.ID(id).FindOne(&shelter)
	if err != nil {
		return nil, err
	} else if !found {
		return nil, errors.New("shelter not found")
	}
	return &shelter, err
}

func (s *NoSQLShelterRepository) FindByName(name string) (*models.Shelter, error) {
	filter := models.Shelter{
		Name: name,
	}
	var shelter models.Shelter
	found, err := s.db.FindOne(&shelter, filter)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("shelter with name %q not found", name)
	}
	return &shelter, nil
}

func (s *NoSQLShelterRepository) FindByLocation(location string) ([]*models.Shelter, error) {
	filter := models.Shelter{
		Location: location,
	}
	var shelter []*models.Shelter
	err := s.db.FindMany(&shelter, filter)
	return shelter, err
}

func (s *NoSQLShelterRepository) FindByOwnerID(id string) ([]*models.Shelter, error) {
	filter := models.Shelter{
		OwnerID: id,
	}
	var shelter []*models.Shelter
	err := s.db.FindMany(&shelter, filter)
	return shelter, err
}

func (s *NoSQLShelterRepository) FindShelters(filter models.Shelter) ([]*models.Shelter, error) {
	s.logger.Infow("finding shelters by filter", "filter", fmt.Sprintf("%+v", filter))
	shelters := make([]*models.Shelter, 0)
	err := s.db.FindMany(&shelters, filter)
	return shelters, err
}

func (s *NoSQLShelterRepository) Save(shelter *models.Shelter) error {
	_, err := s.db.InsertOne(shelter)
	return err
}

func (s *NoSQLShelterRepository) Update(shelter *models.Shelter) error {
	if shelter.ID == "" {
		return errors.New("shelter id missing")
	}

	return s.db.ID(shelter.ID).UpdateOne(shelter)
}

func (s *NoSQLShelterRepository) Delete(id string) error {
	return s.db.ID(id).DeleteOne()
}
