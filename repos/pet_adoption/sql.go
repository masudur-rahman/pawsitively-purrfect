package pet_adoption

import (
	isql "github.com/masudur-rahman/pawsitively-purrfect/infra/database/sql"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/logr"
	"github.com/masudur-rahman/pawsitively-purrfect/models"

	"github.com/rs/xid"
)

type SQLPetAdoptionRepository struct {
	db     isql.Database
	logger logr.Logger
}

func NewSQLPetAdoptionRepository(db isql.Database, logger logr.Logger) *SQLPetAdoptionRepository {
	return &SQLPetAdoptionRepository{
		db:     db.Table("pet_adoption"),
		logger: logger,
	}
}

func (pa *SQLPetAdoptionRepository) AddPetAdoption(petID, userID string) error {
	pa.logger.Infow("adding new pet adoption", "pet", petID, "user", userID)
	adopt := &models.PetAdoption{
		ID:     xid.New().String(),
		UserID: userID,
		PetID:  petID,
	}
	adopt.XKey = adopt.ID
	_, err := pa.db.InsertOne(adopt)
	return err
}

func (pa *SQLPetAdoptionRepository) RemovePetAdoption(petID, userID string) error {
	pa.logger.Infow("removing pet adoption", "pet", petID, "user", userID)
	return pa.db.DeleteOne(models.PetAdoption{
		UserID: userID,
		PetID:  petID,
	})
}

func (pa *SQLPetAdoptionRepository) GetPetOwner(petID string) (string, error) {
	filter := models.PetAdoption{PetID: petID}
	var adopt models.PetAdoption
	has, err := pa.db.FindOne(&adopt, filter)
	if err != nil {
		return "", err
	} else if !has {
		return "", models.ErrPetAdoptionNotFound{PetID: petID}
	}

	return adopt.UserID, nil
}

func (pa *SQLPetAdoptionRepository) ListPetsAdoptedByUser(userID string) ([]string, error) {
	pa.logger.Infow("listing pets by filter", "user", userID)
	filter := models.PetAdoption{UserID: userID}
	adopts := make([]*models.PetAdoption, 0)
	if err := pa.db.FindMany(&adopts, filter); err != nil {
		return nil, err
	}
	pids := make([]string, 0, len(adopts))
	for _, adopt := range adopts {
		pids = append(pids, adopt.PetID)
	}
	return pids, nil
}
