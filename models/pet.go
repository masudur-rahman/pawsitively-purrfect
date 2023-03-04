package models

import "github.com/masudur-rahman/pawsitively-purrfect/models/gqtypes"

type PetAdoptionStatus int

const (
	PetAvailable PetAdoptionStatus = iota
	PetAdopted
)

type Pet struct {
	XKey            string            `json:"_key"`
	ID              string            `json:"id"`
	Name            string            `json:"name"`
	Breed           string            `json:"breed"`
	Gender          string            `json:"gender"`
	Photo           string            `json:"photo"`
	AdoptionStatus  PetAdoptionStatus `json:"adoptionStatus"`
	ShelterID       string            `json:"shelterID""`
	OriginShelterID string            `json:"originShelterID"`
	// If CurrentOwnerID is set, user is the current owner,
	// but it previously belonged to the shelter with the ShelterID
	CurrentOwnerID string `json:"currentOwnerID"`
}

func (s PetAdoptionStatus) String() string {
	switch s {
	case PetAvailable:
		return "Available"
	case PetAdopted:
		return "Adopted"
	}
	return ""
}

func (pet *Pet) APIFormat() gqtypes.Pet {
	return gqtypes.Pet{
		ID:             pet.ID,
		Name:           pet.Name,
		Breed:          pet.Breed,
		Gender:         pet.Gender,
		Photo:          pet.Photo,
		AdoptionStatus: pet.AdoptionStatus.String(),
		ShelterID:      pet.ShelterID,
		CurrentOwnerID: pet.CurrentOwnerID,
	}
}
