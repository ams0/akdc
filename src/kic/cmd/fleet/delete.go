// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"fmt"
	"kic/utils"

	"github.com/spf13/cobra"
)

// deleteCmd deletes a cluster and DNS entry
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an Azure Resource Group and associated Azure DNS record",

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			fmt.Println("Usage: akdc delete serverName")
			return
		}

		fmt.Println("Deleting Resource Group")
		utils.ShellExec(fmt.Sprintf("az group delete -n %s --yes --no-wait", args[0]))

		fmt.Println("Deleting DNS Record")
		utils.ShellExec(fmt.Sprintf("az network dns record-set a delete -g tld -z cseretail.com --yes -n %s", args[0]))

	},
}
