package pet

import (
	"testing"

	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/models/gqtypes"
	petRepo "github.com/masudur-rahman/pawsitively-purrfect/repos/pet"
	userRepo "github.com/masudur-rahman/pawsitively-purrfect/repos/user"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_petService_GetPetByID(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	pr := petRepo.NewMockPetRepository(ctl)
	ps := NewPetService(pr, nil)

	t.Run("pet should not exist", func(t *testing.T) {
		id := "random-id"

		gomock.InOrder(
			pr.EXPECT().FindByID(id).Return(nil, models.ErrPetNotFound{ID: id}),
		)

		pet, err := ps.GetPetByID(id)
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
			pr.EXPECT().FindByID(id).Return(&pf, nil),
		)

		pet, err := ps.GetPetByID(id)
		assert.NoError(t, err)
		assert.NotNil(t, pet)
		assert.EqualValues(t, &pf, pet)
	})
}

func Test_petService_ListShelterPets(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	pr := petRepo.NewMockPetRepository(ctl)
	ps := NewPetService(pr, nil)

	t.Run("pet should not exist", func(t *testing.T) {
		shelterID := "random-shelterID"

		gomock.InOrder(
			pr.EXPECT().FindByShelterID(shelterID).Return([]*models.Pet{}, nil),
		)

		pets, err := ps.ListShelterPets(shelterID)
		assert.NoError(t, err)
		assert.NotNil(t, pets)
		assert.Len(t, pets, 0)
	})

	t.Run("multiple pets found", func(t *testing.T) {
		shelterID := "abc-xyz"

		epets := []*models.Pet{
			&models.Pet{
				ID:        "123",
				Name:      "Cathy",
				Breed:     "Cat",
				Gender:    "Male",
				ShelterID: shelterID,
			},
			&models.Pet{
				ID:        "234",
				Name:      "Cathy",
				Breed:     "Cat",
				Gender:    "Female",
				ShelterID: shelterID,
			},
		}

		gomock.InOrder(
			pr.EXPECT().FindByShelterID(shelterID).Return(epets, nil),
		)

		pets, err := ps.ListShelterPets(shelterID)
		assert.NoError(t, err)
		assert.NotNil(t, pets)
		assert.Len(t, pets, 2)
		assert.Equal(t, epets, pets)
	})
}

func Test_petService_UpdatePet(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	pr := petRepo.NewMockPetRepository(ctl)
	ps := NewPetService(pr, nil)

	t.Run("update pet gender", func(t *testing.T) {
		id := "abc-xyz"

		existingPet := &models.Pet{
			ID:     id,
			Name:   "Cathy",
			Breed:  "Cat",
			Gender: "Male",
		}
		updateParams := gqtypes.PetParams{
			ID:     id,
			Name:   "Cathy",
			Breed:  "Cat",
			Gender: "Female",
		}

		gomock.InOrder(
			pr.EXPECT().FindByID(id).Return(existingPet, nil),
			pr.EXPECT().Update(gomock.Any()).Return(nil),
		)

		pet, err := ps.UpdatePet(updateParams)
		assert.NoError(t, err)
		assert.NotNil(t, pet)
		assert.EqualValues(t, updateParams.Gender, pet.Gender)
	})
}

func Test_petService_AddPetToShelter(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	pr := petRepo.NewMockPetRepository(ctl)
	ps := NewPetService(pr, nil)

	t.Run("update pet gender", func(t *testing.T) {
		params := gqtypes.PetParams{
			Name:      "Cathy",
			Breed:     "Cat",
			Gender:    "Female",
			ShelterID: "1234",
		}

		gomock.InOrder(
			pr.EXPECT().Save(gomock.Any()).Return(nil),
		)

		pet, err := ps.AddPetToShelter(params)
		assert.NoError(t, err)
		assert.NotNil(t, pet)
		assert.EqualValues(t, models.PetAvailable, pet.AdoptionStatus)
	})
}

func Test_petService_AdoptPet(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ur := userRepo.NewMockUserRepository(ctl)
	pr := petRepo.NewMockPetRepository(ctl)
	ps := NewPetService(pr, ur)

	t.Run("update pet gender", func(t *testing.T) {
		userID := "abc-xyz"
		petID := "xyz-abc"

		gomock.InOrder(
			pr.EXPECT().FindByID(petID).Return(&models.Pet{}, nil),
			ur.EXPECT().FindByID(userID).Return(&models.User{}, nil),
			pr.EXPECT().Update(gomock.Any()).Return(nil),
		)

		err := ps.AdoptPet(userID, petID)
		assert.NoError(t, err)
	})
}

func Test_petService_DeletePet(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ur := userRepo.NewMockUserRepository(ctl)
	pr := petRepo.NewMockPetRepository(ctl)
	ps := NewPetService(pr, ur)

	t.Run("update pet gender", func(t *testing.T) {
		id := "abc-xyz"

		gomock.InOrder(
			pr.EXPECT().Delete(gomock.Any()).Return(nil),
		)

		err := ps.DeletePet(id)
		assert.NoError(t, err)
	})
}
