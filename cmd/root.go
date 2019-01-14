// Copyright 2019 Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"

	"github.com/maistra/ior/pkg/galley"
	"github.com/spf13/cobra"
	"istio.io/istio/pkg/log"
)

var (
	loggingOptions = log.DefaultOptions()
	galleyAddr     = "istio-galley.istio-system:9901"
)

func getRootCmd(args []string) *cobra.Command {

	rootCmd := &cobra.Command{
		Use:   "server",
		Short: "Connects to Galley and manages OpenShift Routes based on Istio Gateways",
		Run: func(cmd *cobra.Command, args []string) {
			log.Configure(loggingOptions)
			galley.ConnectToGalley(galleyAddr)
		},
	}

	rootCmd.SetArgs(args)
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	rootCmd.PersistentFlags().StringVarP(&galleyAddr, "mcp-address", "", galleyAddr,
		"Galley's MCP server address")

	loggingOptions.AttachCobraFlags(rootCmd)

	return rootCmd
}
