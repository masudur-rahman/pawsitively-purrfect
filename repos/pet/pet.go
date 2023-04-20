package pet

import (
	"fmt"
	"net/http"

	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/nosql"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/logr"
	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/pkg"

	"github.com/rs/xid"
)

type NoSQLPetRepository struct {
	db     nosql.Database
	logger logr.Logger
}

func NewNoSQLPetRepository(db nosql.Database, logger logr.Logger) *NoSQLPetRepository {
	return &NoSQLPetRepository{
		db:     db.Collection("pet"),
		logger: logger,
	}
}

func (p *NoSQLPetRepository) FindByID(id string) (*models.Pet, error) {
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

func (p *NoSQLPetRepository) FindByIDs(ids []string) ([]*models.Pet, error) {
	query := "FOR doc in pet FILTER doc.id IN @ids RETURN doc"
	bindParams := map[string]interface{}{
		"ids": ids,
	}
	results, err := p.db.Query(query, bindParams)
	if err != nil {
		return nil, err
	}
	var pets []*models.Pet
	if err = pkg.ParseInto(results, &pets); err != nil {
		return nil, err
	}

	return pets, err
}

func (p *NoSQLPetRepository) FindByType(typ models.PetType) ([]*models.Pet, error) {
	filter := models.Pet{
		Type: typ,
	}
	var pets []*models.Pet
	err := p.db.FindMany(&pets, filter)
	return pets, err
}

func (p *NoSQLPetRepository) FindByBreed(breed string) ([]*models.Pet, error) {
	filter := models.Pet{
		Breed: breed,
	}
	var pets []*models.Pet
	err := p.db.FindMany(&pets, filter)
	return pets, err
}

func (p *NoSQLPetRepository) FindByGender(gender string) ([]*models.Pet, error) {
	filter := models.Pet{
		Gender: gender,
	}
	var pets []*models.Pet
	err := p.db.FindMany(&pets, filter)
	return pets, err
}

func (p *NoSQLPetRepository) FindByAdoptionStatus(status models.PetAdoptionStatus) ([]*models.Pet, error) {
	filter := models.Pet{
		AdoptionStatus: status,
	}
	var pets []*models.Pet
	err := p.db.FindMany(&pets, filter)
	return pets, err
}

func (p *NoSQLPetRepository) FindByShelterID(id string) ([]*models.Pet, error) {
	filter := models.Pet{
		ShelterID: id,
	}
	var pets []*models.Pet
	err := p.db.FindMany(&pets, filter)
	return pets, err
}

func (p *NoSQLPetRepository) FindPets(filter models.Pet) ([]*models.Pet, error) {
	p.logger.Infow("finding pets by filter", "filter", fmt.Sprintf("%+v", filter))
	pets := make([]*models.Pet, 0)
	err := p.db.FindMany(&pets, filter)
	return pets, err
}

func (p *NoSQLPetRepository) Save(pet *models.Pet) error {
	p.logger.Infow("adding new pet")
	if pet.ID == "" {
		pet.ID = xid.New().String()
	}
	pet.XKey = pet.ID
	_, err := p.db.InsertOne(pet)
	return err
}

func (p *NoSQLPetRepository) Update(pet *models.Pet) error {
	if pet.ID == "" {
		return models.StatusError{
			Status:  http.StatusBadRequest,
			Message: "pet id missing",
		}
	}

	return p.db.ID(pet.ID).UpdateOne(pet)
}

func (p *NoSQLPetRepository) Delete(id string) error {
	return p.db.ID(id).DeleteOne()
}
