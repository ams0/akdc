// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"github.com/spf13/cobra"
)

// getTokenCmd fetched the admin-user token used to authenticate in the azure portal for kubernetes workloads
var getTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "fetch admin-user token",
	Long:  `fetch admin-user token`,
	Run: func(cmd *cobra.Command, args []string) {
		execClusters("'echo \"$(hostname)\" - $(/home/akdc/get-service-account-token.sh)'", grep)
	},
}
