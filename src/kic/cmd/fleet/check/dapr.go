// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package check

import (
	"github.com/spf13/cobra"
	"kic/utils"
)

// DaprCmd represents the dapr command
var DaprCmd = &cobra.Command{
	Use:   "dapr",
	Short: "Check dapr status on each cluster",

	Args: func(cmd *cobra.Command, args []string) error {
		// this will exit with an error
		utils.ReadHostIPs("")
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		utils.ExecClusters("./gitops/fleet/scripts/check-dapr", grep)
	},
}
