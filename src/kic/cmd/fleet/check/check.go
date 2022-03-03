// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package check

import (
	"github.com/spf13/cobra"
)

// fleetCheckCmd adds check subcommands
var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Check cluster status",
}

func init() {
	// todo - these can all be generated
	CheckCmd.AddCommand(FluxCmd)
	CheckCmd.AddCommand(HeartbeatCmd)
	CheckCmd.AddCommand(LogsCmd)
	CheckCmd.AddCommand(NgsaCmd)
	CheckCmd.AddCommand(RetriesCmd)
	CheckCmd.AddCommand(SetupCmd)
	// CheckCmd.AddCommand(DaprCmd)
	// CheckCmd.AddCommand(RadiusCmd)

	CheckCmd.PersistentFlags().StringVarP(&grep, "grep", "g", "", "grep conditional to filter by host")
}

// option variables
// common across several commands
var grep string
