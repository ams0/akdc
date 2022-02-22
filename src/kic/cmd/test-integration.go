// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"kic/cfmt"
)

// test-integration command
var testIntegrationCmd = &cobra.Command{
	Use:   "integration",
	Short: "run an integration test on each cluster",
	Long:  `run an integration test on each cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		cfmt.Info("Running integration test")

		command := "webv -s http://localhost:30080 "

		command += getTestFlagValues()

		// add test-integration specific options to command line
		if fileIntegration != "" {
			command += " --files ../webv/" + fileIntegration
		}

		if maxErrors > 0 {
			command += fmt.Sprintf(" --max-errors %d", maxErrors)
		}

		if summary != "" {
			command += " --summary " + summary
		}

		shellExec(command)
	},
}

// add command specific options
func init() {
	testIntegrationCmd.Flags().StringVarP(&fileIntegration, "file", "f", "../webv/baseline.json", "Test file to use")
	testIntegrationCmd.Flags().StringVarP(&summary, "summary", "", "None", "Test summary display <None|Tsv|Xml>")
	testIntegrationCmd.Flags().IntVarP(&maxErrors, "max-errors", "", 10, "Max validation errors before terminating test")
	testIntegrationCmd.Flags().IntVarP(&sleep, "sleep", "l", 0, "Sleep (ms) between each request")
}
