package shelter

import (
	"testing"

	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/sql/mock"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/logr"
	"github.com/masudur-rahman/pawsitively-purrfect/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func initializeSqlDatabaseAndShelterRepo(ctl *gomock.Controller) (*mock.MockDatabase, *SQLShelterRepository) {
	db := mock.NewMockDatabase(ctl)
	db.EXPECT().Table("shelter").Return(db).MaxTimes(2)
	ur := NewSQLShelterRepository(db, logr.DefaultLogger)

	return db, ur
}

func TestNewSQLShelterRepository(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, sr := initializeSqlDatabaseAndShelterRepo(ctl)

	assert.NotNil(t, sr.db)
	assert.EqualValues(t, db.Table("shelter"), sr.db)
	assert.Equal(t, logr.DefaultLogger, sr.logger)
}

func TestSQLShelterRepository_FindByID(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, sr := initializeSqlDatabaseAndShelterRepo(ctl)

	t.Run("shelter should not exist", func(t *testing.T) {
		id := "random-id"

		gomock.InOrder(
			db.EXPECT().ID(id).Return(db),
			db.EXPECT().FindOne(gomock.Any(), gomock.Any()).Return(false, models.ErrShelterNotFound{ID: id}),
		)

		user, err := sr.FindByID(id)
		assert.Error(t, err)
		assert.ErrorIs(t, err, models.ErrShelterNotFound{ID: id})
		assert.Nil(t, user)
	})

	t.Run("shelter must exist", func(t *testing.T) {
		id := "abc-xyz"

		sf := models.Shelter{
			ID:          id,
			Name:        "Pawsitive",
			Description: "Pawsitively Purrfect",
			Website:     "http://pawsitively.purrfect",
			Location:    "Bangladesh",
		}

		gomock.InOrder(
			db.EXPECT().ID(id).Return(db),
			db.EXPECT().FindOne(gomock.Any(), gomock.Any()).DoAndReturn(func(shelter *models.Shelter, _ ...interface{}) (bool, error) {
				*shelter = sf
				return true, nil
			}),
		)

		shelter, err := sr.FindByID(id)
		assert.NoError(t, err)
		assert.NotNil(t, shelter)
		assert.EqualValues(t, &sf, shelter)
	})
}

func TestSQLShelterRepository_FindByName(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, sr := initializeSqlDatabaseAndShelterRepo(ctl)

	t.Run("shelter should not exist", func(t *testing.T) {
		name := "random-name"
		filter := models.Shelter{Name: name}
		gomock.InOrder(
			db.EXPECT().FindOne(gomock.Any(), filter).Return(false, models.ErrShelterNotFound{Name: name}),
		)

		shelter, err := sr.FindByName(name)
		assert.Error(t, err)
		assert.ErrorIs(t, err, models.ErrShelterNotFound{Name: name})
		assert.Nil(t, shelter)
	})

	t.Run("shelter must exist", func(t *testing.T) {
		name := "Pawsitive"

		filter := models.Shelter{Name: name}
		sf := models.Shelter{
			ID:          "abc-xyz",
			Name:        "Pawsitive",
			Description: "Pawsitively Purrfect",
			Website:     "http://pawsitively.purrfect",
			Location:    "Bangladesh",
		}

		gomock.InOrder(
			db.EXPECT().FindOne(gomock.Any(), filter).DoAndReturn(func(shelter *models.Shelter, _ ...interface{}) (bool, error) {
				*shelter = sf
				return true, nil
			}),
		)

		shelter, err := sr.FindByName(name)
		assert.NoError(t, err)
		assert.NotNil(t, shelter)
		assert.EqualValues(t, &sf, shelter)
	})
}

func TestSQLShelterRepository_FindByLocation(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, sr := initializeSqlDatabaseAndShelterRepo(ctl)

	t.Run("shelter should not exist", func(t *testing.T) {
		location := "random-location"
		filter := models.Shelter{Location: location}
		gomock.InOrder(
			db.EXPECT().FindMany(gomock.Any(), filter).Return(nil),
		)

		shelters, err := sr.FindByLocation(location)
		assert.Nil(t, err)
		assert.Len(t, shelters, 0)
	})

	t.Run("shelter must exist", func(t *testing.T) {
		location := "Bangladesh"

		filter := models.Shelter{Location: location}
		sf := models.Shelter{
			ID:          "abc-xyz",
			Name:        "Pawsitive",
			Description: "Pawsitively Purrfect",
			Website:     "http://pawsitively.purrfect",
			Location:    "Bangladesh",
		}

		gomock.InOrder(
			db.EXPECT().FindMany(gomock.Any(), filter).DoAndReturn(func(shelters *[]*models.Shelter, filter interface{}) error {
				*shelters = []*models.Shelter{&sf}
				return nil
			}),
		)

		shelters, err := sr.FindByLocation(location)
		assert.NoError(t, err)
		assert.NotNil(t, shelters)
		assert.EqualValues(t, []*models.Shelter{&sf}, shelters)
	})
}

func TestSQLShelterRepository_FindByOwnerID(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, sr := initializeSqlDatabaseAndShelterRepo(ctl)

	t.Run("shelter should not exist", func(t *testing.T) {
		ownerID := "random-ownerID"
		filter := models.Shelter{OwnerID: ownerID}
		gomock.InOrder(
			db.EXPECT().FindMany(gomock.Any(), filter).Return(nil),
		)

		shelters, err := sr.FindByOwnerID(ownerID)
		assert.NoError(t, err)
		assert.Len(t, shelters, 0)
	})

	t.Run("shelter must exist", func(t *testing.T) {
		ownerID := "1"
		filter := models.Shelter{OwnerID: ownerID}
		sf := models.Shelter{
			ID:          "abc-xyz",
			Name:        "Pawsitive",
			Description: "Pawsitively Purrfect",
			Website:     "http://pawsitively.purrfect",
			Location:    "Bangladesh",
			OwnerID:     ownerID,
		}

		gomock.InOrder(
			db.EXPECT().FindMany(gomock.Any(), filter).DoAndReturn(func(shelters *[]*models.Shelter, filter interface{}) error {
				*shelters = []*models.Shelter{&sf}
				return nil
			}),
		)

		shelters, err := sr.FindByOwnerID(ownerID)
		assert.NoError(t, err)
		assert.NotNil(t, shelters)
		assert.EqualValues(t, []*models.Shelter{&sf}, shelters)
	})
}

func TestSQLShelterRepository_FindShelters(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, sr := initializeSqlDatabaseAndShelterRepo(ctl)

	t.Run("no shelters", func(t *testing.T) {
		filter := models.Shelter{
			NumberOfPets: 5,
		}
		gomock.InOrder(
			db.EXPECT().FindMany(gomock.Any(), filter).Return(nil),
		)

		shelters, err := sr.FindShelters(filter)
		assert.NoError(t, err)
		assert.EqualValues(t, 0, len(shelters))
	})
}

func TestSQLShelterRepository_Save(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, sr := initializeSqlDatabaseAndShelterRepo(ctl)

	t.Run("should create shelter", func(t *testing.T) {
		id := "abc-xyz"
		shelter := models.Shelter{
			ID:          id,
			Name:        "Pawsitive",
			Description: "Pawsitively Purrfect",
			Website:     "http://pawsitively.purrfect",
			Location:    "Bangladesh",
		}

		gomock.InOrder(
			db.EXPECT().InsertOne(gomock.Any()).Return(id, nil),
		)

		err := sr.Save(&shelter)
		assert.NoError(t, err)
		assert.Equal(t, id, shelter.XKey)
	})
}

func TestSQLShelterRepository_Update(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, sr := initializeSqlDatabaseAndShelterRepo(ctl)

	t.Run("should update shelter", func(t *testing.T) {
		id := "abc-xyz"
		shelter := models.Shelter{
			ID:          id,
			Name:        "Pawsitive",
			Description: "Pawsitively Purrfect",
			Website:     "http://pawsitively.purrfect",
			Location:    "Bangladesh",
		}

		gomock.InOrder(
			db.EXPECT().ID(id).Return(db),
			db.EXPECT().UpdateOne(gomock.Any()).Return(nil),
		)

		err := sr.Update(&shelter)
		assert.NoError(t, err)
		assert.NotNil(t, shelter)
	})
}

func TestSQLShelterRepository_Delete(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, sr := initializeSqlDatabaseAndShelterRepo(ctl)

	t.Run("should update shelter", func(t *testing.T) {
		id := "abc-xyz"

		gomock.InOrder(
			db.EXPECT().ID(id).Return(db),
			db.EXPECT().DeleteOne().Return(nil),
		)

		err := sr.Delete(id)
		assert.NoError(t, err)
	})
}
