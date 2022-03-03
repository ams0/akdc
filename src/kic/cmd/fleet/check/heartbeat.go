// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package check

import (
	"github.com/spf13/cobra"
	"kic/utils"
)

// HeartbeatCmd checks each cluster for DNS and SSL setup
// tinybench is our "heartbeat" application deployed as part of bootstrap
var HeartbeatCmd = &cobra.Command{
	Use:   "heartbeat",
	Short: "Check https heartbeat on each server",

	Args: func(cmd *cobra.Command, args []string) error {
		// this will exit with an error
		utils.ReadHostIPs("")
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		utils.ExecClusters("./gitops/fleet/scripts/check-heartbeat", grep)
	},
}
