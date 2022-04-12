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
var (
	rootCmd = &cobra.Command{
		Use:   "kic",
		Short: "Retail Edge CLI",
	}
)

// initialize the root command
func init() {
	// load the commands from the bin location ./.appName directory
	boa.LoadCommands(rootCmd)

	// add commands to kic
	if rootCmd.Name() == "kic" {
		addCommandIfNotExist("test", test.TestCmd)
		addCommandIfNotExist("targets", targets.TargetsCmd)
	}

	// add fleet commands
	if rootCmd.Name() == "flt" {
		// add missing commands
		addCommandIfNotExist("create", fleet.CreateCmd)
		addCommandIfNotExist("delete", fleet.DeleteCmd)
		addCommandIfNotExist("exec", fleet.ExecCmd)
		addCommandIfNotExist("list", fleet.ListCmd)
		addCommandIfNotExist("ssh", fleet.SshCmd)
		addCommandIfNotExist("targets", targets.TargetsCmd)

		if check := boa.GetCommandByUse(rootCmd, "check"); check != nil {
			if app := boa.GetCommandByUse(check, "app"); app != nil {
				app.ValidArgsFunction = validArgsFleetCheckApp
				app.Args = cobra.ExactValidArgs(1)
			}
		}
	}

	// this will set a new root if specified
	// this will also remove any hidden commands so they don't exist
	boa.SetNewRoot()
}

// add the command to root if it doesn't exist
func addCommandIfNotExist(name string, cmd *cobra.Command) {
	if boa.GetCommandByUse(rootCmd, name) == nil {
		rootCmd.AddCommand(cmd)
	}
}

// execute the root command
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// validate flt check app arg
func validArgsFleetCheckApp(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	apps, err := boa.ReadCompletionFile("flt-check-app-completion")

	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	// only one arg
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	// sugest from the completion file
	return apps, cobra.ShellCompDirectiveNoFileComp
}
