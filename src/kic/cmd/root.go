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

	// add commands to vm command
	vmCmd := boa.GetCommandByUse(rootCmd, "vms")
	if vmCmd != nil && !vmCmd.Hidden {
		if boa.GetCommandByUse(vmCmd, "test") == nil {
			vmCmd.AddCommand(test.TestCmd)
		}
	}

	// add commands to local command
	localCmd := boa.GetCommandByUse(rootCmd, "local")
	if localCmd != nil && !localCmd.Hidden {
		if boa.GetCommandByUse(localCmd, "targets") == nil {
			localCmd.AddCommand(targets.TargetsCmd)
		}
		if boa.GetCommandByUse(localCmd, "test") == nil {
			localCmd.AddCommand(test.TestCmd)
		}
	}

	// add fleet command to root
	if boa.GetCommandByUse(rootCmd, "fleet") == nil && fleet.FleetCmd != nil {
		rootCmd.AddCommand(fleet.FleetCmd)
		fleet.FleetCmd.AddCommand(targets.TargetsCmd)
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
