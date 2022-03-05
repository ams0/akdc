// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package check

import (
	"kic/boa"

	"github.com/spf13/cobra"
)

// CheckNgsaCmd checks each cluster to see if ngsa is installed
var NgsaCmd = &cobra.Command{
	Use:   "ai-order-accuracy",
	Short: "Check AI Order Accuracy status on each cluster",

	Args: func(cmd *cobra.Command, args []string) error {
		// this will exit with an error
		boa.ReadHostIPs("")

		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		boa.ExecClusters("./gitops/fleet/scripts/check-aioa", grep)
	},
}
