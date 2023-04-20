package pet

import (
	"testing"

	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/sql/mock"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/logr"
	"github.com/masudur-rahman/pawsitively-purrfect/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func initializeSqlDatabaseAndPetRepo(ctl *gomock.Controller) (*mock.MockDatabase, *SQLPetRepository) {
	db := mock.NewMockDatabase(ctl)
	db.EXPECT().Table("pet").Return(db).MaxTimes(2)
	pr := NewSQLPetRepository(db, logr.DefaultLogger)

	return db, pr
}

func TestNewSQLPetRepository(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, pr := initializeSqlDatabaseAndPetRepo(ctl)

	assert.NotNil(t, pr.db)
	assert.EqualValues(t, db.Table("pet"), pr.db)
	assert.Equal(t, logr.DefaultLogger, pr.logger)

}

func TestSQLPetRepository_FindByID(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, pr := initializeSqlDatabaseAndPetRepo(ctl)

	t.Run("pet should not exist", func(t *testing.T) {
		id := "random-id"

		gomock.InOrder(
			db.EXPECT().ID(id).Return(db),
			db.EXPECT().FindOne(gomock.Any(), gomock.Any()).Return(false, models.ErrPetNotFound{ID: id}),
		)

		pet, err := pr.FindByID(id)
		assert.Error(t, err)
		assert.ErrorIs(t, err, models.ErrPetNotFound{ID: id})
		assert.Nil(t, pet)
	})

	t.Run("pet must exist", func(t *testing.T) {
		id := "abc-xyz"

		pf := models.Pet{
			ID:     id,
			Name:   "Cathy",
			Breed:  "Cat",
			Gender: "Male",
		}

		gomock.InOrder(
			db.EXPECT().ID(id).Return(db),
			db.EXPECT().FindOne(gomock.Any(), gomock.Any()).DoAndReturn(func(pet *models.Pet, _ ...interface{}) (bool, error) {
				*pet = pf
				return true, nil
			}),
		)

		pet, err := pr.FindByID(id)
		assert.NoError(t, err)
		assert.NotNil(t, pet)
		assert.EqualValues(t, &pf, pet)
	})
}

func TestSQLPetRepository_FindByName(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, pr := initializeSqlDatabaseAndPetRepo(ctl)

	t.Run("pet should not exist", func(t *testing.T) {
		name := "random-name"
		filter := models.Pet{Name: name}
		gomock.InOrder(
			db.EXPECT().FindOne(gomock.Any(), filter).Return(false, models.ErrPetNotFound{Name: name}),
		)

		pet, err := pr.FindByName(name)
		assert.Error(t, err)
		assert.ErrorIs(t, err, models.ErrPetNotFound{Name: name})
		assert.Nil(t, pet)
	})

	t.Run("pet must exist", func(t *testing.T) {
		name := "Cathy"

		filter := models.Pet{Name: name}
		pf := models.Pet{
			Name:   name,
			Breed:  "Cat",
			Gender: "Male",
		}

		gomock.InOrder(
			db.EXPECT().FindOne(gomock.Any(), filter).DoAndReturn(func(pet *models.Pet, _ ...interface{}) (bool, error) {
				*pet = pf
				return true, nil
			}),
		)

		pet, err := pr.FindByName(name)
		assert.NoError(t, err)
		assert.NotNil(t, pet)
		assert.EqualValues(t, &pf, pet)
	})
}

func TestSQLPetRepository_FindByAdoptionStatus(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, pr := initializeSqlDatabaseAndPetRepo(ctl)

	t.Run("pet should not exist", func(t *testing.T) {
		status := models.PetAvailable

		filter := models.Pet{AdoptionStatus: status}
		gomock.InOrder(
			db.EXPECT().FindMany(gomock.Any(), filter).Return(nil),
		)

		pets, err := pr.FindByAdoptionStatus(status)
		assert.NoError(t, err)
		assert.Len(t, pets, 0)
	})

	t.Run("pet must exist", func(t *testing.T) {
		status := models.PetAvailable

		filter := models.Pet{AdoptionStatus: status}
		pf := models.Pet{
			Name:           "Cathy",
			Breed:          "Cat",
			Gender:         "Male",
			AdoptionStatus: status,
		}

		gomock.InOrder(
			db.EXPECT().FindMany(gomock.Any(), filter).DoAndReturn(func(pets *[]*models.Pet, filter interface{}) error {
				*pets = []*models.Pet{&pf}
				return nil
			}),
		)

		pets, err := pr.FindByAdoptionStatus(status)
		assert.NoError(t, err)
		assert.NotNil(t, pets)
		assert.EqualValues(t, []*models.Pet{&pf}, pets)
	})
}

func TestSQLPetRepository_FindByBreed(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, pr := initializeSqlDatabaseAndPetRepo(ctl)

	t.Run("pet should not exist", func(t *testing.T) {
		breed := "Cat"

		filter := models.Pet{Breed: breed}
		gomock.InOrder(
			db.EXPECT().FindMany(gomock.Any(), filter).Return(nil),
		)

		pets, err := pr.FindByBreed(breed)
		assert.NoError(t, err)
		assert.Len(t, pets, 0)
	})

	t.Run("pet must exist", func(t *testing.T) {
		breed := "Cat"

		filter := models.Pet{Breed: breed}
		pf := models.Pet{
			Name:   "Cathy",
			Breed:  breed,
			Gender: "Male",
		}

		gomock.InOrder(
			db.EXPECT().FindMany(gomock.Any(), filter).DoAndReturn(func(pets *[]*models.Pet, filter interface{}) error {
				*pets = []*models.Pet{&pf}
				return nil
			}),
		)

		pets, err := pr.FindByBreed(breed)
		assert.NoError(t, err)
		assert.NotNil(t, pets)
		assert.EqualValues(t, []*models.Pet{&pf}, pets)
	})
}

func TestSQLPetRepository_FindByGender(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, pr := initializeSqlDatabaseAndPetRepo(ctl)

	t.Run("pet should not exist", func(t *testing.T) {
		gender := "Male"

		filter := models.Pet{Gender: gender}
		gomock.InOrder(
			db.EXPECT().FindMany(gomock.Any(), filter).Return(nil),
		)

		pets, err := pr.FindByGender(gender)
		assert.NoError(t, err)
		assert.Len(t, pets, 0)
	})

	t.Run("pet must exist", func(t *testing.T) {
		gender := "Male"

		filter := models.Pet{Gender: gender}
		pf := models.Pet{
			Name:   "Cathy",
			Gender: gender,
		}

		gomock.InOrder(
			db.EXPECT().FindMany(gomock.Any(), filter).DoAndReturn(func(pets *[]*models.Pet, filter interface{}) error {
				*pets = []*models.Pet{&pf}
				return nil
			}),
		)

		pets, err := pr.FindByGender(gender)
		assert.NoError(t, err)
		assert.NotNil(t, pets)
		assert.EqualValues(t, []*models.Pet{&pf}, pets)
	})
}

func TestSQLPetRepository_FindByShelterID(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, pr := initializeSqlDatabaseAndPetRepo(ctl)

	t.Run("pet should not exist", func(t *testing.T) {
		shelterID := "123"

		filter := models.Pet{ShelterID: shelterID}
		gomock.InOrder(
			db.EXPECT().FindMany(gomock.Any(), filter).Return(nil),
		)

		pets, err := pr.FindByShelterID(shelterID)
		assert.NoError(t, err)
		assert.Len(t, pets, 0)
	})

	t.Run("pet must exist", func(t *testing.T) {
		shelterID := "123"

		filter := models.Pet{ShelterID: shelterID}
		pf := models.Pet{
			Name:      "Cathy",
			Gender:    "Male",
			ShelterID: shelterID,
		}

		gomock.InOrder(
			db.EXPECT().FindMany(gomock.Any(), filter).DoAndReturn(func(pets *[]*models.Pet, filter interface{}) error {
				*pets = []*models.Pet{&pf}
				return nil
			}),
		)

		pets, err := pr.FindByShelterID(shelterID)
		assert.NoError(t, err)
		assert.NotNil(t, pets)
		assert.EqualValues(t, []*models.Pet{&pf}, pets)
	})
}

func TestSQLPetRepository_FindPets(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, pr := initializeSqlDatabaseAndPetRepo(ctl)

	t.Run("pet should not exist", func(t *testing.T) {
		status := models.PetAvailable

		filter := models.Pet{AdoptionStatus: status}
		gomock.InOrder(
			db.EXPECT().FindMany(gomock.Any(), filter).Return(nil),
		)

		pets, err := pr.FindPets(filter)
		assert.NoError(t, err)
		assert.Len(t, pets, 0)
	})

	t.Run("pet must exist", func(t *testing.T) {
		status := models.PetAvailable

		filter := models.Pet{AdoptionStatus: status}
		pf := models.Pet{
			Name:           "Cathy",
			Breed:          "Cat",
			Gender:         "Male",
			AdoptionStatus: status,
		}

		gomock.InOrder(
			db.EXPECT().FindMany(gomock.Any(), filter).DoAndReturn(func(pets *[]*models.Pet, filter interface{}) error {
				*pets = []*models.Pet{&pf}
				return nil
			}),
		)

		pets, err := pr.FindPets(filter)
		assert.NoError(t, err)
		assert.NotNil(t, pets)
		assert.EqualValues(t, []*models.Pet{&pf}, pets)
	})
}

func TestSQLPetRepository_Save(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, pr := initializeSqlDatabaseAndPetRepo(ctl)

	t.Run("should create pet", func(t *testing.T) {
		id := "abc-xyz"
		pet := models.Pet{
			ID:     id,
			Name:   "Cathy",
			Gender: "Male",
		}

		gomock.InOrder(
			db.EXPECT().InsertOne(gomock.Any()).Return(id, nil),
		)

		err := pr.Save(&pet)
		assert.NoError(t, err)
		assert.Equal(t, id, pet.XKey)
	})
}

func TestSQLPetRepository_Update(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, pr := initializeSqlDatabaseAndPetRepo(ctl)

	t.Run("should update pet", func(t *testing.T) {
		id := "abc-xyz"
		pet := models.Pet{
			ID:     id,
			Name:   "Cathy",
			Gender: "Male",
		}

		gomock.InOrder(
			db.EXPECT().ID(id).Return(db),
			db.EXPECT().UpdateOne(gomock.Any()).Return(nil),
		)

		err := pr.Update(&pet)
		assert.NoError(t, err)
		assert.NotNil(t, pet)
	})
}

func TestSQLPetRepository_Delete(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	db, pr := initializeSqlDatabaseAndPetRepo(ctl)

	t.Run("should update pet", func(t *testing.T) {
		id := "abc-xyz"

		gomock.InOrder(
			db.EXPECT().ID(id).Return(db),
			db.EXPECT().DeleteOne().Return(nil),
		)

		err := pr.Delete(id)
		assert.NoError(t, err)
	})
}
