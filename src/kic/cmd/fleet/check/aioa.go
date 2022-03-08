// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package check

import (
	"kic/boa"

	"github.com/spf13/cobra"
)

// CheckAioaCmd checks each cluster to see if AI Order Accuracy is installed
var AioaCmd = &cobra.Command{
	Use:   "ai-order-accuracy",
	Short: "Check AI Order Accuracy status on each cluster",
	Args:  argsFleetCheck,

	RunE: func(cmd *cobra.Command, args []string) error {
		return boa.ExecClusters("./gitops/fleet/scripts/check-aioa", grep)
	},
}
