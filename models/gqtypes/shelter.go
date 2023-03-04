package gqtypes

type Shelter struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Website            string `json:"website"`
	Location           string `json:"location"`
	ContactInformation string `json:"contactInformation"`
	Logo               string `json:"logo"`
	NumberOfPets       int64  `json:"numberOfPets"`
	OwnerID            string `json:"ownerID"`
}

type ShelterParams struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Website            string `json:"website"`
	Location           string `json:"location"`
	ContactInformation string `json:"contactInformation"`
	OwnerID            string `json:"-"`
}
