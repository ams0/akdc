// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// option variables
// common across several commands
var debug bool
var dryRun bool
var region string
var tag string
var zone string

// used by check, exec, sync and test commands
var grep string

// mainly for create command
var cluster string
var group string
var location string
var repo string
var pem string
var key string
var quiet bool
var ssl bool

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

// rootCmd represents the base command and adds commands, options, and flags
var rootCmd = &cobra.Command{
	Use:   "akdc",
	Short: "Retail Edge CLI",
	Long:  `Retail Edge CLI`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// initialize the root command
func init() {
	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(execCmd)
	rootCmd.AddCommand(groupsCmd)
	rootCmd.AddCommand(sshCmd)
	rootCmd.AddCommand(syncCmd)
	rootCmd.AddCommand(targetsCmd)
	rootCmd.AddCommand(testCmd)
}
