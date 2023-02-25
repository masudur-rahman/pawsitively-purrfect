package models

import "github.com/masudur-rahman/pawsitively-purrfect/models/gqtypes"

type Shelter struct {
	XKey               string `json:"_key"`
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Website            string `json:"website"`
	Location           string `json:"location"`
	ContactInformation string `json:"contactInformation"`
	Logo               string `json:"logo"`
	NumberOfPets       int64  `json:"numberOfPets"`
	// Assuming a shelter can have only one owner
	OwnerID string `json:"ownerID"`
}

func (s *Shelter) APIFormat() gqtypes.Shelter {
	return gqtypes.Shelter{
		ID:                 s.ID,
		Name:               s.Name,
		Description:        s.Description,
		Website:            s.Website,
		Location:           s.Location,
		ContactInformation: s.ContactInformation,
		Logo:               s.Logo,
		NumberOfPets:       s.NumberOfPets,
		OwnerID:            s.OwnerID,
	}
}
