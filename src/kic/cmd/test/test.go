// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package test

import (
	"fmt"
	"github.com/spf13/cobra"
)

// mainly for test commands
var dryRun bool
var logFormat string
var region string
var sleep int
var tag string
var verbose bool
var zone string

// test parent command
var TestCmd = &cobra.Command{
	Use:   "test",
	Short: "Run cluster tests",
}

// initialize the parent command
func init() {
	// add sub-commands
	TestCmd.AddCommand(IntegrationCmd)
	TestCmd.AddCommand(LoadCmd)

	// add common options
	TestCmd.PersistentFlags().BoolVarP(&dryRun, "dry-run", "d", false, "Validate configuration without running")
	TestCmd.PersistentFlags().StringVarP(&logFormat, "log-format", "", "", "Log format <Json|JsonCamel|None|Tsv|TsvMin>")
	TestCmd.PersistentFlags().StringVarP(&region, "region", "", "", "Region deployed to (user defined)")
	TestCmd.PersistentFlags().IntVarP(&sleep, "sleep", "l", 0, "Sleep (ms) between each request")
	TestCmd.PersistentFlags().StringVarP(&tag, "tag", "", "", "Tag for log (user defined)")
	TestCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Show verbose results")
	TestCmd.PersistentFlags().StringVarP(&zone, "zone", "", "", "Zone deployed to (user defined)")
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
