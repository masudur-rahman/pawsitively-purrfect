package cmd

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/masudur-rahman/pawsitively-purrfect/configs"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/nosql"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/nosql/arangodb"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/logr"
	"github.com/masudur-rahman/pawsitively-purrfect/pkg"
	"github.com/masudur-rahman/pawsitively-purrfect/services/all"

	"gopkg.in/yaml.v3"
)

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
}

func initialize(ctx context.Context) *all.Services {
	switch configs.PurrfectConfig.Database.Type {
	case configs.DatabaseArangoDB:
		return getServicesForArangoDB(ctx)
	case configs.DatabasePostgres:
		return getServicesForPostgres(ctx)
	default:
		return nil
	}
}

func getServicesForArangoDB(ctx context.Context) *all.Services {
	arangoDB, err := arangodb.InitializeArangoDB(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	var db nosql.Database
	db = arangodb.NewArangoDB(ctx, arangoDB)
	logger := logr.DefaultLogger
	return all.GetNoSQLServices(db, logger)
}

func getServicesForPostgres(ctx context.Context) *all.Services {
	panic("return sql services")
}
