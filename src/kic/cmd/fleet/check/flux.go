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
	Args:  argsFleetCheck,

	RunE: func(cmd *cobra.Command, args []string) error {
		return boa.ExecClusters("./fleet-vm/scripts/check-flux", grep)
	},
}
