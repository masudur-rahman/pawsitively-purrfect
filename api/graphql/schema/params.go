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

var updateProfileFieldArgs = graphql.FieldConfigArgument{
	"id":        &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.String}},
	"username":  &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.String}},
	"email":     &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.String}},
	"firstName": &graphql.ArgumentConfig{Type: graphql.String},
	"lastName":  &graphql.ArgumentConfig{Type: graphql.String},
	"bio":       &graphql.ArgumentConfig{Type: graphql.String},
	"location":  &graphql.ArgumentConfig{Type: graphql.String},
	"avatar":    &graphql.ArgumentConfig{Type: graphql.String},
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

var adoptionStatusEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "AdoptionStatus",
	Values: graphql.EnumValueConfigMap{
		"Available": &graphql.EnumValueConfig{
			Value: "Available",
		},
		"Adopted": &graphql.EnumValueConfig{
			Value: "Adopted",
		},
	},
})

var listShelterPetFieldArgs = graphql.FieldConfigArgument{
	"shelterID":      &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.ID}},
	"adoptionStatus": &graphql.ArgumentConfig{Type: adoptionStatusEnum},
}

var findPetsFieldArgs = graphql.FieldConfigArgument{
	"name":           &graphql.ArgumentConfig{Type: graphql.String},
	"type":           &graphql.ArgumentConfig{Type: petTypeEnum},
	"breed":          &graphql.ArgumentConfig{Type: graphql.String},
	"gender":         &graphql.ArgumentConfig{Type: genderEnum},
	"shelterID":      &graphql.ArgumentConfig{Type: graphql.String},
	"adoptionStatus": &graphql.ArgumentConfig{Type: adoptionStatusEnum},
}

var petTypeEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "PetType",
	Values: graphql.EnumValueConfigMap{
		"Cat": &graphql.EnumValueConfig{
			Value: "Cat",
		},
		"Dog": &graphql.EnumValueConfig{
			Value: "Dog",
		},
	},
})

var genderEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "Gender",
	Values: graphql.EnumValueConfigMap{
		"Male": &graphql.EnumValueConfig{
			Value: "Male",
		},
		"Female": &graphql.EnumValueConfig{
			Value: "Female",
		},
	},
})

var addPetFieldArgs = graphql.FieldConfigArgument{
	"name":      &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.String}},
	"type":      &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: petTypeEnum}},
	"breed":     &graphql.ArgumentConfig{Type: graphql.String},
	"gender":    &graphql.ArgumentConfig{Type: genderEnum},
	"shelterID": &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.String}},
}

var updatePetFieldArgs = graphql.FieldConfigArgument{
	"id":     &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.String}},
	"name":   &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.String}},
	"type":   &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: petTypeEnum}},
	"breed":  &graphql.ArgumentConfig{Type: graphql.String},
	"gender": &graphql.ArgumentConfig{Type: genderEnum},
	//"shelterID": &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.String}},
}
