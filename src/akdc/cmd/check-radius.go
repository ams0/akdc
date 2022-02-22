// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"github.com/spf13/cobra"
)

// checkRadiusCmd checks each cluster to verify radius is installed
var checkRadiusCmd = &cobra.Command{
	Use:   "radius",
	Short: "check radius status on each cluster",
	Long:  `check radius status on each cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		execClusters("'if [[ $(kubectl get ns) == *\"radius-system\"* ]]; then echo \"$(hostname) success\"; else echo \"$(hostname) failed\"; fi'", grep)
	},
}
