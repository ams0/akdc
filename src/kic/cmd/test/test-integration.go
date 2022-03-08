// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package test

import (
	"fmt"
	"kic/boa"
	"kic/boa/cfmt"
	"os"

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
		Short: "Run an integration test on each cluster",
		RunE:  runTestIntegrationE,
	}
)

// add command specific options
func init() {
	IntegrationCmd.Flags().StringVarP(&fileIntegration, "file", "f", "baseline.json", "Test file to use")
	IntegrationCmd.Flags().IntVarP(&maxErrors, "max-errors", "", 10, "Max validation errors before terminating test")
	IntegrationCmd.Flags().StringVarP(&summary, "summary", "", "None", "Test summary display <None|Tsv|Xml>")
}

// run the test-integration command
func runTestIntegrationE(cmd *cobra.Command, args []string) error {
	cfmt.Info("Running integration test")

	command := "test"

	if sleep < 0 {
		sleep = 0
	}

	// get shared options
	params := getTestFlagValues()

	// add test-integration specific options to command line
	if fileIntegration != "" {
		params += " --files " + fileIntegration
	}

	if maxErrors > 0 {
		params += fmt.Sprintf(" --max-errors %d", maxErrors)
	}

	if summary != "" {
		params += " --summary " + summary
	}

	path := boa.GetBinDir() + "/.kic/commands/" + command

	// execute the file with "bash -c" if it exists
	if _, err := os.Stat(path); err == nil {
		return boa.ShellExecE(fmt.Sprintf("%s %s", path, params))
	} else {
		return cfmt.ErrorE(err)
	}
}
