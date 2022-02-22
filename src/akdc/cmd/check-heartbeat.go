// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"github.com/spf13/cobra"
)

// checkSslCmd checks each cluster for DNS and SSL setup
// tinybench is our "heartbeat" application deployed as part of bootstrap
var checkHeartbeatCmd = &cobra.Command{
	Use:   "heartbeat",
	Short: "check https heartbeat on each server",
	Long:  `check https heartbeat on each server`,
	Run: func(cmd *cobra.Command, args []string) {
		execClusters("'echo \"$(curl -s https://$(hostname).cseretail.com/heartbeat/17)  $(hostname)\"'", grep)
	},
}
