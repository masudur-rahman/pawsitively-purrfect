package shelter

import (
	"fmt"
	"net/http"

	isql "github.com/masudur-rahman/pawsitively-purrfect/infra/database/sql"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/logr"
	"github.com/masudur-rahman/pawsitively-purrfect/models"

	"github.com/rs/xid"
)

type SQLShelterRepository struct {
	db     isql.Database
	logger logr.Logger
}

func NewSQLShelterRepository(db isql.Database, logger logr.Logger) *SQLShelterRepository {
	return &SQLShelterRepository{
		db:     db.Table("shelter"),
		logger: logger,
	}
}

func (s *SQLShelterRepository) FindByID(id string) (*models.Shelter, error) {
	var shelter models.Shelter
	found, err := s.db.ID(id).FindOne(&shelter)
	if err != nil {
		return nil, err
	} else if !found {
		return nil, models.ErrShelterNotFound{ID: id}
	}
	return &shelter, err
}

func (s *SQLShelterRepository) FindByName(name string) (*models.Shelter, error) {
	filter := models.Shelter{
		Name: name,
	}
	var shelter models.Shelter
	found, err := s.db.FindOne(&shelter, filter)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, models.ErrShelterNotFound{Name: name}
	}
	return &shelter, nil
}

func (s *SQLShelterRepository) FindByLocation(location string) ([]*models.Shelter, error) {
	filter := models.Shelter{
		Location: location,
	}
	var shelter []*models.Shelter
	err := s.db.FindMany(&shelter, filter)
	return shelter, err
}

func (s *SQLShelterRepository) FindByOwnerID(id string) ([]*models.Shelter, error) {
	filter := models.Shelter{
		OwnerID: id,
	}
	var shelter []*models.Shelter
	err := s.db.FindMany(&shelter, filter)
	return shelter, err
}

func (s *SQLShelterRepository) FindShelters(filter models.Shelter) ([]*models.Shelter, error) {
	s.logger.Infow("finding shelters by filter", "filter", fmt.Sprintf("%+v", filter))
	shelters := make([]*models.Shelter, 0)
	err := s.db.FindMany(&shelters, filter)
	return shelters, err
}

func (s *SQLShelterRepository) Save(shelter *models.Shelter) error {
	s.logger.Infow("creating new shelter")
	if shelter.ID == "" {
		shelter.ID = xid.New().String()
	}
	shelter.XKey = shelter.ID
	_, err := s.db.InsertOne(shelter)
	return err
}

func (s *SQLShelterRepository) Update(shelter *models.Shelter) error {
	if shelter.ID == "" {
		return models.StatusError{
			Status:  http.StatusBadRequest,
			Message: "shelter id missing",
		}
	}

	return s.db.ID(shelter.ID).UpdateOne(shelter)
}

func (s *SQLShelterRepository) Delete(id string) error {
	return s.db.ID(id).DeleteOne()
}
