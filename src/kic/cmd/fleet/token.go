// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"github.com/spf13/cobra"
	"kic/utils"
)

// getTokenCmd fetched the admin-user token used to authenticate in the azure portal for kubernetes workloads
var TokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Get Arc token from each cluster",

	Run: func(cmd *cobra.Command, args []string) {
		utils.ExecClusters("'echo -e \"\"$(hostname)\"\n$(/home/akdc/get-service-account-token.sh)\n\"'", grep)
	},
}

func init() {
	TokenCmd.Flags().StringVarP(&grep, "grep", "g", "", "grep conditional to filter by host")
}
