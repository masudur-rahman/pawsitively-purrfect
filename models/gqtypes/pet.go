package gqtypes

type Pet struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Breed          string `json:"breed,omitempty"`
	Gender         string `json:"gender"`
	Photo          string `json:"photo"`
	AdoptionStatus string `json:"adoptionStatus"`
	ShelterID      string `json:"shelterID,omitempty"`
}

type PetParams struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Breed          string `json:"breed"`
	Gender         string `json:"gender"`
	ShelterID      string `json:"shelterID"`
	AdoptionStatus string `json:"adoptionStatus"`
}
