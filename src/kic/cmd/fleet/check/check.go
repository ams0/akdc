// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package check

import (
	"kic/boa"

	"github.com/spf13/cobra"
)

var (
	// option variables
	grep string

	// fleetCheckCmd adds check subcommands
	CheckCmd = &cobra.Command{
		Use:   "check",
		Short: "Check cluster status",
	}
)

func init() {
	// todo - these can all be generated
	CheckCmd.AddCommand(FluxCmd)
	CheckCmd.AddCommand(HeartbeatCmd)
	CheckCmd.AddCommand(LogsCmd)
	CheckCmd.AddCommand(AioaCmd)
	CheckCmd.AddCommand(RetriesCmd)
	CheckCmd.AddCommand(SetupCmd)
	CheckCmd.AddCommand(TLogCmd)
	// CheckCmd.AddCommand(DaprCmd)
	// CheckCmd.AddCommand(RadiusCmd)

	CheckCmd.PersistentFlags().StringVarP(&grep, "grep", "g", "", "grep conditional to filter by host")
}

// check the args
func argsFleetCheck(cmd *cobra.Command, args []string) error {
	// this will exit with an error
	boa.ReadHostIPs("")
	return nil
}
