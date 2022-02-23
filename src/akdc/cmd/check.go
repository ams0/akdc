// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"github.com/spf13/cobra"
)

// checkCmd adds check subcommands
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check cluster status",
	Long:  `check cluster status`,
}

func init() {
	checkCmd.AddCommand(checkDaprCmd)
	checkCmd.AddCommand(checkFluxCmd)
	checkCmd.AddCommand(checkHeartbeatCmd)
	checkCmd.AddCommand(checkLogsCmd)
	checkCmd.AddCommand(checkNgsaCmd)
	checkCmd.AddCommand(checkRadiusCmd)
	checkCmd.AddCommand(checkRetriesCmd)
	checkCmd.AddCommand(checkSetupCmd)
	checkCmd.PersistentFlags().String("async", "a", "run checks asynchronously")

	checkCmd.PersistentFlags().StringVarP(&grep, "grep", "g", "", "grep conditional to filter by host")
}
