package repos

type PetAdoptionRepository interface {
	AddPetAdoption(petID, userID string) error
	RemovePetAdoption(petID, userID string) error
	GetPetOwner(petID string) (string, error)
	ListPetsAdoptedByUser(userID string) ([]string, error)
}
