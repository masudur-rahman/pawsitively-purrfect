package pet

import (
	"fmt"
	"net/http"

	isql "github.com/masudur-rahman/pawsitively-purrfect/infra/database/sql"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/logr"
	"github.com/masudur-rahman/pawsitively-purrfect/models"

	"github.com/rs/xid"
)

type SQLPetRepository struct {
	db     isql.Database
	logger logr.Logger
}

func NewSQLPetRepository(db isql.Database, logger logr.Logger) *SQLPetRepository {
	return &SQLPetRepository{
		db:     db.Table("pet"),
		logger: logger,
	}
}

func (p *SQLPetRepository) FindByID(id string) (*models.Pet, error) {
	var pet models.Pet
	found, err := p.db.ID(id).FindOne(&pet)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, models.ErrPetNotFound{ID: id}
	}
	return &pet, nil
}

func (p *SQLPetRepository) FindByIDs(ids []string) ([]*models.Pet, error) {
	filter := map[string]interface{}{
		"id": ids,
	}
	var pets []*models.Pet
	err := p.db.FindMany(&pets, filter)
	return pets, err
}

func (p *SQLPetRepository) FindByType(typ models.PetType) ([]*models.Pet, error) {
	filter := models.Pet{
		Type: typ,
	}
	var pets []*models.Pet
	err := p.db.FindMany(&pets, filter)
	return pets, err
}

func (p *SQLPetRepository) FindByBreed(breed string) ([]*models.Pet, error) {
	filter := models.Pet{
		Breed: breed,
	}
	var pets []*models.Pet
	err := p.db.FindMany(&pets, filter)
	return pets, err
}

func (p *SQLPetRepository) FindByGender(gender string) ([]*models.Pet, error) {
	filter := models.Pet{
		Gender: gender,
	}
	var pets []*models.Pet
	err := p.db.FindMany(&pets, filter)
	return pets, err
}

func (p *SQLPetRepository) FindByAdoptionStatus(status models.PetAdoptionStatus) ([]*models.Pet, error) {
	filter := models.Pet{
		AdoptionStatus: status,
	}
	var pets []*models.Pet
	err := p.db.FindMany(&pets, filter)
	return pets, err
}

func (p *SQLPetRepository) FindByShelterID(id string) ([]*models.Pet, error) {
	filter := models.Pet{
		ShelterID: id,
	}
	var pets []*models.Pet
	err := p.db.FindMany(&pets, filter)
	return pets, err
}

func (p *SQLPetRepository) FindPets(filter models.Pet) ([]*models.Pet, error) {
	p.logger.Infow("finding pets by filter", "filter", fmt.Sprintf("%+v", filter))
	pets := make([]*models.Pet, 0)
	err := p.db.FindMany(&pets, filter)
	return pets, err
}

func (p *SQLPetRepository) Save(pet *models.Pet) error {
	p.logger.Infow("adding new pet")
	if pet.ID == "" {
		pet.ID = xid.New().String()
	}
	pet.XKey = pet.ID
	_, err := p.db.InsertOne(pet)
	return err
}

func (p *SQLPetRepository) Update(pet *models.Pet) error {
	if pet.ID == "" {
		return models.StatusError{
			Status:  http.StatusBadRequest,
			Message: "pet id missing",
		}
	}

	return p.db.ID(pet.ID).UpdateOne(pet)
}

func (p *SQLPetRepository) Delete(id string) error {
	return p.db.ID(id).DeleteOne()
}
