// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package test

import (
	"fmt"
	"kic/boa"
	"kic/boa/cfmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	// command line options
	fileIntegration string
	maxErrors       int
	summary         string

	// test-integration command
	IntegrationCmd = &cobra.Command{
		Use:   "integration",
		Short: "Run an integration test on the cluster",
		RunE:  runTestIntegrationE,
	}
)

// add command specific options
func init() {
	IntegrationCmd.Flags().StringVarP(&fileIntegration, "file", "f", "", "Test file to use")
	IntegrationCmd.Flags().IntVarP(&maxErrors, "max-errors", "", 10, "Max validation errors before terminating test")
	IntegrationCmd.Flags().StringVarP(&summary, "summary", "", "None", "Test summary display <None|Tsv|Xml>")
}

// run the test-integration command
func runTestIntegrationE(cmd *cobra.Command, args []string) error {
	cfmt.Info("Running integration test")

	if sleep < 0 {
		sleep = 0
	}

	// get shared options
	params := getTestFlagValues()

	if maxErrors > 0 {
		params += fmt.Sprintf(" --max-errors %d", maxErrors)
	}

	if summary != "" {
		params += " --summary " + summary
	}

	if fileIntegration == "" {
		if boa.GetBinName() == "kivm" {
			fileIntegration = "imdb-baseline.json heartbeat-baseline.json "
		} else {
			fileIntegration = "imdb-baseline.json "
		}
	}

	// add test-integration specific options to command line
	if fileIntegration != "" {
		params += " --files " + fileIntegration
	}

	// get the webv container
	webv := os.Getenv("AKDC_WEBV")

	if webv == "" {
		webv = "ghcr.io/cse-labs/webv-red:latest"
	}

	// build the path to the script
	path := "docker run --net host --rm " + webv + " --server "

	if boa.GetBinName() == "kivm" {
		path += "http://$AKDC_FQDN "
	} else {
		path += "http://localhost:30080 "
	}

	path += " " + params

	if len(args) > 0 {
		path += " " + strings.Join(args, " ")
	}

	// execute the file with "bash -c" if it exists
	return boa.ShellExecE(path)
}
