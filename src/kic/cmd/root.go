// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"kic/boa"
	"kic/cmd/fleet"
	"kic/cmd/fleet/check"
	"kic/cmd/targets"
	"kic/cmd/test"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command and adds commands, options, and flags
var rootCmd = &cobra.Command{
	Use:   "akdc",
	Short: "Retail Edge CLI",
}

// initialize the root command
func init() {
	// load the commands from the bin location ./.appName directory
	boa.LoadCommands(rootCmd)

	// add commands to kic
	if rootCmd.Name() == "kic" {

		if boa.GetCommandByUse(rootCmd, "test") == nil {
			rootCmd.AddCommand(test.TestCmd)
		}

		if boa.GetCommandByUse(rootCmd, "targets") == nil {
			rootCmd.AddCommand(targets.TargetsCmd)
		}

	}

	// add fleet commands
	if rootCmd.Name() == "flt" {
		rootCmd.AddCommand(check.CheckCmd)
		rootCmd.AddCommand(fleet.CreateCmd)
		rootCmd.AddCommand(fleet.DeleteCmd)
		rootCmd.AddCommand(fleet.ExecCmd)
		rootCmd.AddCommand(fleet.GroupsCmd)
		rootCmd.AddCommand(fleet.ListCmd)
		rootCmd.AddCommand(fleet.PatchCmd)
		rootCmd.AddCommand(fleet.PullCmd)
		rootCmd.AddCommand(fleet.SshCmd)
		rootCmd.AddCommand(fleet.SyncCmd)
		rootCmd.AddCommand(fleet.ArcTokenCmd)
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
