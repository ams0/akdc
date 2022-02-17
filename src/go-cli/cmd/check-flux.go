// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"github.com/spf13/cobra"
)

// checkFluxCmd checks each cluster for flux-check namespace
var checkFluxCmd = &cobra.Command{
	Use:   "flux",
	Short: "check flux status on each cluster",
	Long:  `check flux status on each cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		execClusters("'if [[ $(kubectl get ns) == *\"flux-check\"* ]]; then echo \"$(hostname) success\"; else echo \"$(hostname) failed\"; fi'", grep)
	},
}
