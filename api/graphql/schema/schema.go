package schema

import (
	"github.com/masudur-rahman/pawsitively-purrfect/api/graphql/resolvers"

	"github.com/graphql-go/graphql"
)

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.ID},
		"username":  &graphql.Field{Type: graphql.String},
		"email":     &graphql.Field{Type: graphql.String},
		"firstName": &graphql.Field{Type: graphql.String},
		"lastName":  &graphql.Field{Type: graphql.String},
		"bio":       &graphql.Field{Type: graphql.String},
		"location":  &graphql.Field{Type: graphql.String},
		"avatar":    &graphql.Field{Type: graphql.String},
		"isActive":  &graphql.Field{Type: graphql.Boolean},
		"isAdmin":   &graphql.Field{Type: graphql.Boolean},
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
		"shelterID":      &graphql.Field{Type: graphql.ID},
		"currentOwnerID": &graphql.Field{Type: graphql.ID},
	},
})

var shelterType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Shelter",
	Fields: graphql.Fields{
		"id":                 &graphql.Field{Type: graphql.ID},
		"name":               &graphql.Field{Type: graphql.String},
		"description":        &graphql.Field{Type: graphql.String},
		"website":            &graphql.Field{Type: graphql.String},
		"location":           &graphql.Field{Type: graphql.String},
		"contactInformation": &graphql.Field{Type: graphql.String},
		"logo":               &graphql.Field{Type: graphql.String},
		"numberOfPets":       &graphql.Field{Type: graphql.Int},
		"ownerID":            &graphql.Field{Type: graphql.ID},
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
					"id": &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.ID}},
				},
				Resolve: resolver.GetShelter,
			},
			"listShelters": &graphql.Field{
				Type:        graphql.NewList(shelterType),
				Description: "List shelters by filter",
				Args:        shelterFilterFieldArgs,
				Resolve:     resolver.ListShelters,
			},

			"pet": &graphql.Field{
				Type:        petType,
				Description: "Get a pet by ID.",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.ID}},
				},
				Resolve: resolver.GetPet,
			},
			"listPets": &graphql.Field{
				Type:        graphql.NewList(petType),
				Description: "List pets by shelter",
				Args: graphql.FieldConfigArgument{
					"shelterID": &graphql.ArgumentConfig{Type: &graphql.NonNull{OfType: graphql.ID}},
				},
				Resolve: resolver.ListPets,
			},
		},
	})

	return query
}

func rootMutation(resolver *resolvers.Resolver) *graphql.Object {
	mutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			// user
			"register": &graphql.Field{
				Type:        userType,
				Description: "Register a new user to the system",
				Args:        registerParams,
				Resolve:     resolver.RegisterUser,
			},

			"login": &graphql.Field{
				Type:        userType,
				Description: "Login user to the system",
				Args:        loginParams,
				Resolve:     resolver.Login,
			},

			"updateProfile": &graphql.Field{
				Type:        userType,
				Description: "Update user profile",
				Args:        updateProfileFieldArgs,
				Resolve:     resolver.UpdateProfile,
			},

			// shelter
			"addShelter": &graphql.Field{
				Type:        shelterType,
				Description: "Add new shelter to the system",
				Args:        addShelterFieldArgs,
				Resolve:     resolver.AddShelter,
			},
			"updateShelter": &graphql.Field{
				Type:        shelterType,
				Description: "Update shelter information",
				Args:        updateShelterFieldArgs,
				Resolve:     nil, // TODO
			},

			"addPet": &graphql.Field{
				Type:        petType,
				Description: "Add new pet to a shelter",
				Args:        addPetFieldArgs,
				Resolve:     resolver.AddPetNewPet,
			},

			"updatePet": &graphql.Field{
				Type:        petType,
				Description: "Update pet information",
				Args:        updatePetFieldArgs,
				Resolve:     nil, // TODO
			},

			"adoptPet": &graphql.Field{
				Type:        petType,
				Description: "Adopt pet from a shelter",
				Args: graphql.FieldConfigArgument{
					"petID": &graphql.ArgumentConfig{Type: graphql.ID},
				},
				Resolve: nil, // TODO
			},
		},
	})

	return mutation
}

func PurrfectSchema(resolver *resolvers.Resolver) (graphql.Schema, error) {
	query := rootQuery(resolver)
	mutation := rootMutation(resolver)
	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	})
}
