// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// mainly for test commands
var verbose bool
var fileIntegration string
var fileLoad string
var duration int
var sleep int
var port int
var random bool
var logFormat string
var summary string
var maxErrors int
var delayStart int
var dryRun bool
var region string
var zone string
var tag string

// test parent command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "run cluster tests",
	Long:  `run cluster tests`,
}

// initialize the parent command
func init() {
	// add sub-commands
	testCmd.AddCommand(testIntegrationCmd)
	testCmd.AddCommand(testLoadCmd)

	// add common options
	testCmd.PersistentFlags().BoolVarP(&dryRun, "dry-run", "d", false, "Validate configuration without running")
	testCmd.PersistentFlags().StringVarP(&region, "region", "", "", "Region deployed to (user defined)")
	testCmd.PersistentFlags().StringVarP(&tag, "tag", "", "", "Tag for log (user defined)")
	testCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Show verbose results")
	testCmd.PersistentFlags().StringVarP(&zone, "zone", "", "", "Zone deployed to (user defined)")
	testCmd.PersistentFlags().StringVarP(&logFormat, "log-format", "", "", "Log format <Json|JsonCamel|None|Tsv|TsvMin>")
}

// add the flags to the command line
func getTestFlagValues() string {
	cmd := ""

	if verbose {
		cmd += " --verbose"
	}
	if sleep > 0 {
		cmd += fmt.Sprintf(" --sleep %d", sleep)
	}
	if dryRun {
		cmd += " --dry-run"
	}
	if region != "" {
		cmd += "--region " + region
	}
	if zone != "" {
		cmd += "--zone " + zone
	}
	if tag != "" {
		cmd += " --tag " + tag
	}
	if logFormat != "" {
		cmd += " --log-format " + logFormat
	}

	return cmd
}
