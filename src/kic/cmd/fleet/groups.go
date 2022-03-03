// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"kic/utils"

	"github.com/spf13/cobra"
)

// GroupsCmd gets the Azure Resource Groups from the subscription
var GroupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "List Azure resource groups",

	Run: func(cmd *cobra.Command, args []string) {
		utils.ShellExec("az group list -o table | sort | grep -e central- -e east- -e west- -e corp- -e retail-edge")
	},
}
