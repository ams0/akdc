// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"kic/boa"

	"github.com/spf13/cobra"
)

// PatchCmd runs ~/fleet-vm/scripts/patch.sh on each cluster
var PatchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Run a patch command on each cluster",
	RunE: func(cmd *cobra.Command, args []string) error {
		return (boa.ExecClusters("./fleet-vm/scripts/patch.sh", grep))
	},
}

func init() {
	PatchCmd.Flags().StringVarP(&grep, "grep", "g", "", "grep conditional to filter by host")
}
