package schema

import "github.com/graphql-go/graphql"

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.ID,
			Description: "The ID of the user.",
		},
		"username": &graphql.Field{
			Type:        graphql.String,
			Description: "The username of the user.",
		},
		"email": &graphql.Field{
			Type:        graphql.String,
			Description: "The email of the user.",
		},
		"fullName": &graphql.Field{
			Type:        graphql.String,
			Description: "Full name of the user",
		},
		"bio": &graphql.Field{
			Type:        graphql.String,
			Description: "The bio of the user.",
		},
		"location": &graphql.Field{
			Type:        graphql.String,
			Description: "The location of the user.",
		},
		"avatar": &graphql.Field{
			Type:        graphql.String,
			Description: "The avatar of the user.",
		},
		"isActive": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Whether the user is active or not.",
		},
		"isAdmin": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Whether the user is admin or not.",
		},
	},
})

var petType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Pet",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.ID,
			Description: "The ID of the pet.",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "The name of the pet.",
		},
		"breed": &graphql.Field{
			Type:        graphql.String,
			Description: "The breed of the pet.",
		},
		"gender": &graphql.Field{
			Type:        graphql.String,
			Description: "The gender of the pet.",
		},
		"photo": &graphql.Field{
			Type:        graphql.String,
			Description: "URL of the pet's photo.",
		},
		"adoptionStatus": &graphql.Field{
			Type:        graphql.String,
			Description: "The adoption status of the pet.",
		},
		"shelterID": &graphql.Field{
			Type:        graphql.String,
			Description: "The shelter ID of the pet.",
		},
		"currentOwnerID": &graphql.Field{
			Type:        graphql.String,
			Description: "The current owner ID of the pet.",
		},
	},
})

var shelterType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Shelter",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.ID,
			Description: "The ID of the shelter.",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "The name of the shelter.",
		},
		"website": &graphql.Field{
			Type:        graphql.String,
			Description: "Website for the shelter",
		},
		"location": &graphql.Field{
			Type:        graphql.String,
			Description: "The location of the shelter.",
		},
		"contactInformation": &graphql.Field{
			Type:        graphql.String,
			Description: "The contact information of the shelter.",
		},
		"description": &graphql.Field{
			Type:        graphql.String,
			Description: "The description of the shelter.",
		},
		"logo": &graphql.Field{
			Type:        graphql.String,
			Description: "The logo url of the shelter.",
		},
		"numberOfPets": &graphql.Field{
			Type:        graphql.Int,
			Description: "The number of pets of the shelter.",
		},
		"ownerID": &graphql.Field{
			Type:        graphql.Int,
			Description: "The owner ID of the shelter.",
		},
	},
})
