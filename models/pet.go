package models

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

type PetAdoptionStatus string

const (
	PetAdopted   = "adopted"
	PetAvailable = "available"
)
