// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"github.com/spf13/cobra"
)

// checkSetupCmd checks each cluster for the setup status
var checkSetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "check setup status on each cluster",
	Long:  `check setup status on each cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		execClusters("'echo \"$(hostname) $(tail -n1 status)\"'", grep)
	},
}
