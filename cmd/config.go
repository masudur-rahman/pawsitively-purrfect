package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/masudur-rahman/pawsitively-purrfect/api/graphql/resolvers"
	"github.com/masudur-rahman/pawsitively-purrfect/configs"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/nosql"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/nosql/arangodb"
	"github.com/masudur-rahman/pawsitively-purrfect/pkg"
	"github.com/masudur-rahman/pawsitively-purrfect/repos/pet"
	"github.com/masudur-rahman/pawsitively-purrfect/repos/shelter"
	"github.com/masudur-rahman/pawsitively-purrfect/repos/user"
	petsvc "github.com/masudur-rahman/pawsitively-purrfect/services/pet"
	sheltersvc "github.com/masudur-rahman/pawsitively-purrfect/services/shelter"
	usersvc "github.com/masudur-rahman/pawsitively-purrfect/services/user"

	"github.com/go-logr/logr"
	"github.com/the-redback/go-oneliners"
	"gopkg.in/yaml.v3"
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

	fmt.Println(shelterRepo, petRepo)

	userSvc := usersvc.NewUserService(userRepo)
	shelterSvc := sheltersvc.NewShelterService(shelterRepo)
	petSvc := petsvc.NewPetService(petRepo, userRepo)

	return resolvers.NewResolver(userSvc, shelterSvc, petSvc)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile == "" {
		cfgFile = filepath.Join(pkg.ProjectDirectory, "configs", ".pawsitively-purrfect.yaml")
	}

	data, err := os.ReadFile(cfgFile)
	if err != nil {
		log.Fatalf("Reading config file %v, %v", cfgFile, err)
	}

	if err = yaml.Unmarshal(data, &configs.PurrfectConfig); err != nil {
		log.Fatalf("Unmarshaling PurrfectConfig, %v", err)
	}

	oneliners.PrettyJson(configs.PurrfectConfig, "Purrfect Config")
}
