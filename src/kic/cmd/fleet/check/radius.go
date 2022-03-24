// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package check

import (
	"kic/boa"

	"github.com/spf13/cobra"
)

// RadiusCmd checks each cluster to verify radius is installed
var RadiusCmd = &cobra.Command{
	Use:   "radius",
	Short: "Check radius status on each cluster",
	Args:  argsFleetCheck,

	RunE: func(cmd *cobra.Command, args []string) error {
		return boa.ExecClusters("./fleet-vm/scripts/check-radius", grep)
	},
}
