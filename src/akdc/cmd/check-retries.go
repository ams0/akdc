// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"github.com/spf13/cobra"
)

// checkRetriesCmd checks each server for the number of flux retries during installation
var checkRetriesCmd = &cobra.Command{
	Use:   "retries",
	Short: "check number of retries on each cluster",
	Long:  `check number of retries on each cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		execClusters("'echo \"$(hostname) $(cat status | grep -e retries | tail -1)\"'", grep)
	},
}
