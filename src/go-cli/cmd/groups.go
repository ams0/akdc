// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"github.com/spf13/cobra"
)

// groupsCmd gets the Azure Resource Groups from the subscription
var groupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "list Azure resource groups",
	Long: `list Azure resource groups`,
	Run: func(cmd *cobra.Command, args []string) {
		shellExec("az group list -o table | sort | grep -e central- -e east- -e west- -e corp- -e retail-edge")
	},
}
