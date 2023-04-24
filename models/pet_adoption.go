package models

type PetAdoption struct {
	XKey   string `json:"_key"`
	ID     string `json:"id"`
	PetID  string `json:"petID"`
	UserID string `json:"userID"`
}
