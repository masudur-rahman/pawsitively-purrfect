package schema

import (
	"github.com/masudur-rahman/pawsitively-purrfect/api/graphql/resolvers"

	"github.com/graphql-go/graphql"
)

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":       &graphql.Field{Type: graphql.ID},
		"username": &graphql.Field{Type: graphql.String},
		"email":    &graphql.Field{Type: graphql.String},
		"fullName": &graphql.Field{Type: graphql.String},
		"bio":      &graphql.Field{Type: graphql.String},
		"location": &graphql.Field{Type: graphql.String},
		"avatar":   &graphql.Field{Type: graphql.String},
		"isActive": &graphql.Field{Type: graphql.Boolean},
		"isAdmin":  &graphql.Field{Type: graphql.Boolean},
	},
})

var petType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Pet",
	Fields: graphql.Fields{
		"id":             &graphql.Field{Type: graphql.ID},
		"name":           &graphql.Field{Type: graphql.String},
		"breed":          &graphql.Field{Type: graphql.String},
		"gender":         &graphql.Field{Type: graphql.String},
		"photo":          &graphql.Field{Type: graphql.String},
		"adoptionStatus": &graphql.Field{Type: graphql.String},
		"shelterID":      &graphql.Field{Type: graphql.String},
		"currentOwnerID": &graphql.Field{Type: graphql.String},
	},
})

var shelterType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Shelter",
	Fields: graphql.Fields{
		"id":                 &graphql.Field{Type: graphql.ID},
		"name":               &graphql.Field{Type: graphql.String},
		"website":            &graphql.Field{Type: graphql.String},
		"location":           &graphql.Field{Type: graphql.String},
		"contactInformation": &graphql.Field{Type: graphql.String},
		"description":        &graphql.Field{Type: graphql.String},
		"logo":               &graphql.Field{Type: graphql.String},
		"numberOfPets":       &graphql.Field{Type: graphql.Int},
		"ownerID":            &graphql.Field{Type: graphql.Int},
	},
})

func rootQuery(resolver *resolvers.Resolver) *graphql.Object {
	query := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type:        userType,
				Description: "Get a user by ID.",
				Args: graphql.FieldConfigArgument{
					"id":   &graphql.ArgumentConfig{Type: graphql.ID},
					"name": &graphql.ArgumentConfig{Type: graphql.String},
				},
				Resolve: resolver.GetUser,
			},
			"shelter": &graphql.Field{
				Type:        shelterType,
				Description: "Get a shelter by ID.",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.ID},
				},
				Resolve: resolver.GetShelter,
			},
			"pet": &graphql.Field{
				Type:        petType,
				Description: "Get a pet by ID.",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.ID},
				},
				Resolve: resolver.GetPet,
			},
		},
	})

	return query
}

func PurrfectSchema(resolver *resolvers.Resolver) (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery(resolver),
	})
}
