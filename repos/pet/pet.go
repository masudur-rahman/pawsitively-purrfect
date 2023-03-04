package pet

import (
	"errors"
	"fmt"

	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/nosql"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/logr"
	"github.com/masudur-rahman/pawsitively-purrfect/models"

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

func (p *NoSQLPetRepository) FindByName(name string) (*models.Pet, error) {
	filter := models.Pet{
		Name: name,
	}
	var pet models.Pet
	found, err := p.db.FindOne(&pet, filter)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, models.ErrPetNotFound{Name: name}
	}
	return &pet, nil
}

func (p *NoSQLPetRepository) FindByCurrentOwnerID(ownerID string) ([]*models.Pet, error) {
	filter := models.Pet{
		CurrentOwnerID: ownerID,
	}
	var pets []*models.Pet
	err := p.db.FindMany(&pets, filter)
	if err != nil {
		return nil, err
	}
	return pets, nil
}

func (p *NoSQLPetRepository) FindByOriginShelterID(id string) ([]*models.Pet, error) {
	filter := models.Pet{
		OriginShelterID: id,
	}
	var pets []*models.Pet
	err := p.db.FindMany(&pets, filter)
	if err != nil {
		return nil, err
	}
	return pets, nil
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
		return errors.New("pet id missing")
	}

	return p.db.ID(pet.ID).UpdateOne(pet)
}

func (p *NoSQLPetRepository) Delete(id string) error {
	return p.db.ID(id).DeleteOne()
}
