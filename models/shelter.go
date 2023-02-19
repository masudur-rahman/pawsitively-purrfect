package models

type Shelter struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Website            string `json:"website"`
	Location           string `json:"location"`
	ContactInformation string `json:"contactInformation"`
	Description        string `json:"description"`
	Logo               string `json:"logo"`
	NumberOfPets       int64  `json:"numberOfPets"`
	// Assuming a shelter can have only one owner
	OwnerID string `json:"ownerID"`
}
