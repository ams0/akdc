// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package check

import (
	"kic/boa"

	"github.com/spf13/cobra"
)

// FluxCmd checks each cluster for flux-check namespace
var FluxCmd = &cobra.Command{
	Use:   "flux",
	Short: "Check flux status on each cluster",

	Args: func(cmd *cobra.Command, args []string) error {
		// this will exit with an error
		boa.ReadHostIPs("")
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		boa.ExecClusters("./gitops/fleet/scripts/check-flux", grep)
	},
}
