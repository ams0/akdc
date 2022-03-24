// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"kic/boa"

	"github.com/spf13/cobra"
)

// TokenCmd displays the admin-user token used to authenticate in the Arc portal
var ArcTokenCmd = &cobra.Command{
	Use:   "arc-token",
	Short: "Get Arc token from each cluster",

	RunE: func(cmd *cobra.Command, args []string) error {
		return (boa.ExecClusters("'echo -e \"\"$(hostname)\"\n$(./fleet-vm/scripts/get-service-account-token.sh)\n\"'", grep))
	},
}

func init() {
	ArcTokenCmd.Flags().StringVarP(&grep, "grep", "g", "", "grep conditional to filter by host")
}
