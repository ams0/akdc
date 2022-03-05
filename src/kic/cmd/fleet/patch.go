// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"kic/boa"

	"github.com/spf13/cobra"
)

// PatchCmd runs ~/gitops/fleet/scripts/patch.sh on each cluster
var PatchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Run a patch command on each cluster",
	Run: func(cmd *cobra.Command, args []string) {
		boa.ExecClusters("gitops/fleet/scripts/patch.sh", grep)
	},
}

func init() {
	PatchCmd.Flags().StringVarP(&grep, "grep", "g", "", "grep conditional to filter by host")
}
