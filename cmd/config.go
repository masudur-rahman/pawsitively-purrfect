package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/masudur-rahman/pawsitively-purrfect/api/graphql/resolvers"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/nosql"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/nosql/arangodb"
	"github.com/masudur-rahman/pawsitively-purrfect/repos/pet"
	"github.com/masudur-rahman/pawsitively-purrfect/repos/shelter"
	"github.com/masudur-rahman/pawsitively-purrfect/repos/user"
	"github.com/masudur-rahman/pawsitively-purrfect/services"

	"github.com/go-logr/logr"
)

func initialize(ctx context.Context) *resolvers.Resolver {
	arangoDB, err := arangodb.InitializeArangoDB(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	var db nosql.Database
	db = arangodb.NewArangoDB(ctx, arangoDB)
	userRepo := user.NewNoSQLUserRepository(db, logr.New(nil))
	shelterRepo := shelter.NewNoSQLShelterRepository(db, logr.New(nil))
	petRepo := pet.NewNoSQLPetRepository(db, logr.New(nil))

	fmt.Println(userRepo, shelterRepo, petRepo)

	var userSvc services.UserService
	var shelterSvc services.ShelterService
	var petSvc services.PetService

	return resolvers.NewResolver(userSvc, shelterSvc, petSvc)
}
