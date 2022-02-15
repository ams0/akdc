// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"github.com/spf13/cobra"
)

// checkNgsaCmd checks each cluster to see if ngsa is installed
var checkNgsaCmd = &cobra.Command{
	Use:   "ngsa",
	Short: "check ngsa status on each cluster",
	Long:  `check ngsa status on each cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		execClusters("'if [[ $(kubectl get pods -A) == *\"ngsa\"* ]]; then echo \"$(hostname) found\"; else echo \"$(hostname) not found\"; fi'")
	},
}
