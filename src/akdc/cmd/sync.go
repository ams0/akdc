// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"github.com/spf13/cobra"
)

// syncCmd runs flux sync (reconcile) on each cluster
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "sync (reconcile) flux on each cluster",
	Long:  `sync (reconcile) flux on each cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		execClusters("flux reconcile source git gitops", grep)
	},
}

func init() {
	syncCmd.Flags().StringVarP(&grep, "grep", "g", "", "grep conditional to filter by host")
}
