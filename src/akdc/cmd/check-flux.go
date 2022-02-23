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
		command := `'
# verify all 3 kustomizations are setup
if [ $(flux get kustomizations | grep -e flux-system -e apps -e bootstrap | wc -l) = 3 ] &&
   # verify all kustomizations are Ready
   [ $(flux get kustomizations | grep -e flux-system -e apps -e bootstrap | cut -f2 | grep False |wc -l) = 0 ]
then
	echo "$(hostname) success"
else
	echo "$(hostname) failed"
fi
'`
		execClusters(command, grep)
	},
}
