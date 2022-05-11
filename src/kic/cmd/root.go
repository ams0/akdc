// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"fmt"
	"kic/boa"
	"kic/boa/cfmt"
	"kic/cmd/fleet"
	"kic/cmd/kic"
	"kic/cmd/kivm"
	"kic/cmd/kubekic"
	"os"

	"github.com/spf13/cobra"
)

var (
	// set TargetCli [and Version] in build via -ldflags - see Makefile
	TargetCli = "flt"
	Version   = "0.4.0"

	// rootCmd represents the base command
	rootCmd = &cobra.Command{}
)

// initialize the root command
func init() {

	// load the commands based on target
	switch TargetCli {
	case "kic":
		rootCmd = kic.AddCommands()
	case "kubekic":
		rootCmd = kubekic.AddCommands()
	case "kivm":
		rootCmd = kivm.AddCommands()
	case "flt":
		rootCmd = fleet.LoadCommands()
	default:
		cfmt.ErrorE("unknown CLI")
		os.Exit(1)
	}

	// add version command
	boa.AddScriptCommand(rootCmd, "version", TargetCli+" version", fmt.Sprintf("echo \"%s\"", Version))
}

// execute the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
