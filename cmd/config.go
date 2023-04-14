package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/masudur-rahman/pawsitively-purrfect/configs"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/nosql"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/nosql/arangodb"
	isql "github.com/masudur-rahman/pawsitively-purrfect/infra/database/sql"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/sql/postgres"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/sql/postgres/pb"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/logr"
	"github.com/masudur-rahman/pawsitively-purrfect/pkg"
	"github.com/masudur-rahman/pawsitively-purrfect/services/all"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	serverAddr := fmt.Sprintf("%s:%v", configs.PurrfectConfig.GRPC.ClientHost, configs.PurrfectConfig.GRPC.Port)
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	//defer conn.Close()

	client := pb.NewPostgresClient(conn)

	var db isql.Database
	db = postgres.NewDatabase(ctx, client)

	logger := logr.DefaultLogger
	return all.GetSQLServices(db, logger)
}
