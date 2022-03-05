// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package check

import (
	"kic/boa"

	"github.com/spf13/cobra"
)

// LogsCmd represents the dapr command
var LogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Check the cloudinit logs on the VMs",

	Args: func(cmd *cobra.Command, args []string) error {
		// this will exit with an error
		boa.ReadHostIPs("")
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		boa.ExecClusters("./gitops/fleet/scripts/check-logs", grep)
	},
}
