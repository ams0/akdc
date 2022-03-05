// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"kic/boa"
	"kic/cmd/fleet"
	"kic/cmd/targets"
	"kic/cmd/test"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command and adds commands, options, and flags
var rootCmd = &cobra.Command{
	Use:   "kic",
	Short: "Kubernetes in Codespaces CLI",
}

// initialize the root command
func init() {
	// load the commands from the bin location ./.appName directory
	boa.LoadCommands(rootCmd)

	// add built-in commands if not exits
	if boa.GetCommandByUse(rootCmd, "fleet") == nil {
		// this can't be a boa command [easily]
		// this is the old akdc set of commands
		rootCmd.AddCommand(fleet.FleetCmd)
	}

	if boa.GetCommandByUse(rootCmd, "test") == nil {
		// this can't be a boa command [easily]
		//    because of the rich command line params
		rootCmd.AddCommand(test.TestCmd)
	}

	if boa.GetCommandByUse(rootCmd, "targets") == nil {
		// this can't be a boa command [easily]
		//    modifies autogitops.json
		rootCmd.AddCommand(targets.TargetsCmd)
	}

	// this will set a new root if specified
	// this will also remove any hidden commands so they don't exist
	boa.SetNewRoot()
}

// execute the root command
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
