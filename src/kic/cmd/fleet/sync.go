// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"kic/boa"

	"github.com/spf13/cobra"
)

// syncCmd runs flux sync (reconcile) on each cluster
var SyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync (reconcile) flux on each cluster",

	Run: func(cmd *cobra.Command, args []string) {
		boa.ExecClusters("flux reconcile source git gitops", grep)
	},
}

func init() {
	SyncCmd.Flags().StringVarP(&grep, "grep", "g", "", "grep conditional to filter by host")
}
