// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package check

import (
	"kic/boa"

	"github.com/spf13/cobra"
)

// HeartbeatCmd checks each cluster for DNS and SSL setup
// tinybench is our "heartbeat" application deployed as part of bootstrap
var HeartbeatCmd = &cobra.Command{
	Use:   "heartbeat",
	Short: "Check https heartbeat on each server",
	Args:  argsFleetCheck,

	RunE: func(cmd *cobra.Command, args []string) error {
		return boa.ExecClusters("./fleet-vm/scripts/check-heartbeat", grep)
	},
}
