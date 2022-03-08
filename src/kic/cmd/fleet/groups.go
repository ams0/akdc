// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"kic/boa"

	"github.com/spf13/cobra"
)

// GroupsCmd gets the Azure Resource Groups from the subscription
var GroupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "List Azure resource groups",

	RunE: func(cmd *cobra.Command, args []string) error {
		return (boa.ShellExecE("az group list -o table | sort | grep -e central- -e east- -e west- -e corp- -e retail-edge"))
	},
}
