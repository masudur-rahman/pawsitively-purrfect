package pet_adoption

import (
	"testing"

	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/sql/mock"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/logr"
	"github.com/masudur-rahman/pawsitively-purrfect/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func initializeSqlDatabaseAndPetAdoptionRepo(ctl *gomock.Controller) (*mock.MockDatabase, *SQLPetAdoptionRepository) {
	db := mock.NewMockDatabase(ctl)
	db.EXPECT().Table("pet_adoption").Return(db).MaxTimes(2)
	pr := NewSQLPetAdoptionRepository(db, logr.DefaultLogger)

	return db, pr
}

func TestNewSQLPetAdoptionRepository(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, par := initializeSqlDatabaseAndPetAdoptionRepo(ctl)

	assert.NotNil(t, par.db)
	assert.EqualValues(t, db.Table("pet_adoption"), par.db)
	assert.Equal(t, logr.DefaultLogger, par.logger)

}

func TestSQLPetAdoptionRepository_AddPetAdoption(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, par := initializeSqlDatabaseAndPetAdoptionRepo(ctl)

	t.Run("add new pet adoption", func(t *testing.T) {
		petID := "123"
		userID := "456"

		gomock.InOrder(
			db.EXPECT().InsertOne(gomock.Any()).Return("", nil),
		)

		err := par.AddPetAdoption(petID, userID)
		assert.NoError(t, err)
	})
}

func TestSQLPetAdoptionRepository_GetPetOwner(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, par := initializeSqlDatabaseAndPetAdoptionRepo(ctl)

	t.Run("no pet owner", func(t *testing.T) {
		petID := "123"

		gomock.InOrder(
			db.EXPECT().FindOne(gomock.Any(), gomock.Any()).Return(false, nil),
		)

		id, err := par.GetPetOwner(petID)
		assert.Error(t, err)
		assert.ErrorIs(t, err, models.ErrPetAdoptionNotFound{PetID: petID})
		assert.EqualValues(t, "", id)
	})

	t.Run("get pet owner", func(t *testing.T) {
		petID := "123"
		ownerID := "abcd"

		gomock.InOrder(
			db.EXPECT().FindOne(gomock.Any(), gomock.Any()).DoAndReturn(func(pa *models.PetAdoption, _ ...interface{}) (bool, error) {
				pa.PetID = petID
				pa.UserID = ownerID
				return true, nil
			}),
		)

		id, err := par.GetPetOwner(petID)
		assert.NoError(t, err)
		assert.EqualValues(t, ownerID, id)
	})
}

func TestSQLPetAdoptionRepository_ListPetsAdoptedByUser(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, par := initializeSqlDatabaseAndPetAdoptionRepo(ctl)

	t.Run("zero pet adoption list", func(t *testing.T) {
		userID := "123"
		filter := models.PetAdoption{UserID: userID}

		gomock.InOrder(
			db.EXPECT().FindMany(gomock.Any(), filter).Return(nil),
		)

		ids, err := par.ListPetsAdoptedByUser(userID)
		assert.NoError(t, err)
		assert.Len(t, ids, 0)
	})

	t.Run("pet adoption list with multiple items", func(t *testing.T) {
		userID := "123"
		filter := models.PetAdoption{UserID: userID}
		pas := []*models.PetAdoption{
			{
				PetID:  "1",
				UserID: "123",
			},
			{
				PetID:  "2",
				UserID: "123",
			},
		}

		gomock.InOrder(
			db.EXPECT().FindMany(gomock.Any(), filter).DoAndReturn(func(pa *[]*models.PetAdoption, filter interface{}) error {
				*pa = pas
				return nil
			}),
		)

		ids, err := par.ListPetsAdoptedByUser(userID)
		assert.NoError(t, err)
		assert.Len(t, ids, 2)
	})
}

func TestSQLPetAdoptionRepository_RemovePetAdoption(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, par := initializeSqlDatabaseAndPetAdoptionRepo(ctl)

	t.Run("remove pet adoption", func(t *testing.T) {
		petID := "123"
		userID := "456"

		gomock.InOrder(
			db.EXPECT().DeleteOne(gomock.Any()).Return(nil),
		)

		err := par.RemovePetAdoption(petID, userID)
		assert.NoError(t, err)
	})
}
