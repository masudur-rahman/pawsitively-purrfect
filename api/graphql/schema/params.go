package schema

import "github.com/graphql-go/graphql"

var registerParams = graphql.FieldConfigArgument{
	"username": &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.String}},
	"email":    &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.String}},
	"password": &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.String}},
}

var loginParams = graphql.FieldConfigArgument{
	"username": &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.String}},
	"password": &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.String}},
}

var shelterFilterFieldArgs = graphql.FieldConfigArgument{
	"name":        &graphql.ArgumentConfig{Type: graphql.String},
	"description": &graphql.ArgumentConfig{Type: graphql.String},
	"website":     &graphql.ArgumentConfig{Type: graphql.String},
	"location":    &graphql.ArgumentConfig{Type: graphql.String},
	"ownerID":     &graphql.ArgumentConfig{Type: graphql.ID},
}

var addShelterFieldArgs = graphql.FieldConfigArgument{
	"name":               &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.String}},
	"description":        &graphql.ArgumentConfig{Type: graphql.String},
	"website":            &graphql.ArgumentConfig{Type: graphql.String},
	"location":           &graphql.ArgumentConfig{Type: graphql.String},
	"contactInformation": &graphql.ArgumentConfig{Type: graphql.String},
}

var updateShelterFieldArgs = graphql.FieldConfigArgument{
	"id":                 &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.String}},
	"name":               &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.String}},
	"description":        &graphql.ArgumentConfig{Type: graphql.String},
	"website":            &graphql.ArgumentConfig{Type: graphql.String},
	"location":           &graphql.ArgumentConfig{Type: graphql.String},
	"contactInformation": &graphql.ArgumentConfig{Type: graphql.String},
}

var addPetFieldArgs = graphql.FieldConfigArgument{
	"name":      &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.String}},
	"breed":     &graphql.ArgumentConfig{Type: graphql.String},
	"gender":    &graphql.ArgumentConfig{Type: graphql.String},
	"shelterID": &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.String}},
}

var updatePetFieldArgs = graphql.FieldConfigArgument{
	"id":        &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.String}},
	"name":      &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.String}},
	"breed":     &graphql.ArgumentConfig{Type: graphql.String},
	"gender":    &graphql.ArgumentConfig{Type: graphql.String},
	"shelterID": &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.String}},
}
