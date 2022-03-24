// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package check

import (
	"kic/boa"

	"github.com/spf13/cobra"
)

// DaprCmd represents the dapr command
var DaprCmd = &cobra.Command{
	Use:   "dapr",
	Short: "Check dapr status on each cluster",
	Args:  argsFleetCheck,

	RunE: func(cmd *cobra.Command, args []string) error {
		return boa.ExecClusters("./fleet-vm/scripts/check-dapr", grep)
	},
}
