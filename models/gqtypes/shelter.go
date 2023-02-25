package gqtypes

import "github.com/graphql-go/graphql"

type ShelterParams struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Website            string `json:"website"`
	Location           string `json:"location"`
	ContactInformation string `json:"contactInformation"`
	OwnerID            string `json:"-"`
}

var AddShelterFieldArgs = graphql.FieldConfigArgument{
	"name":               &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.String}},
	"description":        &graphql.ArgumentConfig{Type: graphql.String},
	"website":            &graphql.ArgumentConfig{Type: graphql.String},
	"location":           &graphql.ArgumentConfig{Type: graphql.String},
	"contactInformation": &graphql.ArgumentConfig{Type: graphql.String},
}

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
