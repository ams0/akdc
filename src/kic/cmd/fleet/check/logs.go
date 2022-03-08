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
	Args:  argsFleetCheck,

	RunE: func(cmd *cobra.Command, args []string) error {
		return boa.ExecClusters("./gitops/fleet/scripts/check-logs", grep)
	},
}
