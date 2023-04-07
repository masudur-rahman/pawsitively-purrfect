/*
Copyright © 2022 Masudur Rahman <masudjuly02@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"

	"github.com/masudur-rahman/pawsitively-purrfect/configs"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/sql/postgres/server"

	"github.com/spf13/cobra"
)

// grpcCmd represents the grpc-serve command
var grpcCmd = &cobra.Command{
	Use:     "grpc-serve",
	Short:   "Start the Postgresql gRPC Server",
	Example: "pawsitively-purrfect grpc-serve",
	Run:     runGRPCServer,
}

func runGRPCServer(cmd *cobra.Command, args []string) {
	if configs.PurrfectConfig.Database.Type != configs.DatabasePostgres {
		log.Fatalln("gRPC only enabled for postgres database")
	}

	if err := server.StartPostgresServer("", "8080"); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	rootCmd.AddCommand(grpcCmd)
}
