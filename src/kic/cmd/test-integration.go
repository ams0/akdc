// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"kic/cfmt"
	"os"
)

// test-integration command
var testIntegrationCmd = &cobra.Command{
	Use:   "integration",
	Short: "run an integration test on each cluster",

	Run: func(cmd *cobra.Command, args []string) {
		cfmt.Info("Running integration test")

		command := "test"

		if sleep < 0 {
			sleep = 0
		}

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

		path := getParentDir() + "/.kic/commands/" + command

		// execute the file with "bash -c" if it exists
		if _, err := os.Stat(path); err == nil {
			shellExec(fmt.Sprintf("%s %s", path, params))
		} else {
			cfmt.Error(err)
		}
	},
}

// add command specific options
func init() {
	testIntegrationCmd.Flags().StringVarP(&fileIntegration, "file", "f", "baseline.json", "Test file to use")
	testIntegrationCmd.Flags().StringVarP(&summary, "summary", "", "None", "Test summary display <None|Tsv|Xml>")
	testIntegrationCmd.Flags().IntVarP(&maxErrors, "max-errors", "", 10, "Max validation errors before terminating test")
}
