package models

type Shelter struct {
	ID                 string
	Name               string
	Website            string
	Location           string
	ContactInformation string
	Description        string
	Logo               string
	NumberOfPets       int64
	// Assuming a shelter can have only one owner
	OwnerID string
}
