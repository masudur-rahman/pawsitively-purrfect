package gqtypes

type Pet struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Breed          string `json:"breed"`
	Gender         string `json:"gender"`
	Photo          string `json:"photo"`
	AdoptionStatus string `json:"adoptionStatus"`
	ShelterID      string `json:"shelterID,omitempty"`
	CurrentOwnerID string `json:"currentOwnerID,omitempty"`
}

type PetParams struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Breed     string `json:"breed"`
	Gender    string `json:"gender"`
	ShelterID string `json:"shelterID"`
}
