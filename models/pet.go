package models

type Pet struct {
	ID             string
	Name           string
	Breed          string
	Gender         string
	Photo          string
	AdoptionStatus string
	ShelterID      int64
	// If CurrentOwnerID is set, user is the current owner,
	// but it previously belonged to the shelter with the ShelterID
	CurrentOwnerID int64
}
