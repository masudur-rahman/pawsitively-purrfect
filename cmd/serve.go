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
	"github.com/masudur-rahman/pawsitively-purrfect/api/http"
	"github.com/masudur-rahman/pawsitively-purrfect/configs"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:     "serve",
	Short:   "Start the Pawsitively Purrfect GraphQL Server",
	Example: "pawsitively-purrfect serve",
	Run:     runServe,
}

func runServe(cmd *cobra.Command, args []string) {
	scfg := configs.PurrfectConfig.Server
	svc := initialize(cmd.Context())

	f := http.Routes(svc)
	f.Run(scfg.Host, scfg.Port)
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
