// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"github.com/spf13/cobra"
)

// daprCmd represents the dapr command
var checkDaprCmd = &cobra.Command{
	Use:   "dapr",
	Short: "check dapr status on each cluster",
	Long:  `check dapr status on each cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		execClusters("'if [[ $(kubectl get ns) == *\"dapr-system\"* ]]; then echo \"$(hostname) success\"; else echo \"$(hostname) failed\"; fi'", grep)
	},
}
