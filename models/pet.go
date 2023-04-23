package models

import "github.com/masudur-rahman/pawsitively-purrfect/models/gqtypes"

type PetAdoptionStatus int

const (
	PetAvailable PetAdoptionStatus = iota
	PetAdopted
)

type PetType string

const (
	PetCat PetType = "Cat"
	PetDog PetType = "Dog"
)

type Pet struct {
	XKey           string            `json:"_key"`
	ID             string            `json:"id"`
	Name           string            `json:"name"`
	Type           PetType           `json:"type"`
	Breed          string            `json:"breed"`
	Gender         string            `json:"gender"`
	Photo          string            `json:"photo"`
	AdoptionStatus PetAdoptionStatus `json:"adoptionStatus"`

	// ShelterID will always be set, even if any user adopts it
	// AdoptionStatus will be set to adopted, if anyone adopts it
	ShelterID string `json:"shelterID"`
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
		Type:           string(pet.Type),
		Breed:          pet.Breed,
		Gender:         pet.Gender,
		Photo:          pet.Photo,
		AdoptionStatus: pet.AdoptionStatus.String(),
		ShelterID:      pet.ShelterID,
	}
}
