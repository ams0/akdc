// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"fmt"
	"kic/boa"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DeleteCmd deletes a cluster and DNS entry
var DeleteCmd = &cobra.Command{
	Use:               "delete",
	Short:             "Delete an Azure Resource Group and associated Azure DNS record",
	Args:              cobra.ExactValidArgs(1),
	ValidArgsFunction: validArgsFleetDelete,
	RunE:              runFleetDeleteE,
}

// run kic fleet delete command
func runFleetDeleteE(cmd *cobra.Command, args []string) error {
	// check if resource group exists
	res, _ := boa.ShellExecOut(fmt.Sprintf("az group exists -n %s", args[0]), false)

	if strings.TrimSpace(res) == "true" {
		fmt.Println("Deleting Resource Group")
		boa.ShellExecE(fmt.Sprintf("az group delete -g %s --yes --no-wait", args[0]))
	}

	// delete the DNS record (if exists)
	boa.ShellExecE("flt dns delete " + args[0])

	return nil
}

// validate kic fleet delete args
func validArgsFleetDelete(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	// don't use the defaultIPs
	if _, err := os.Stat("ips"); err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	// only one argument
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	// suggest from the ips file
	if ips, err := boa.ReadHostIPs(""); err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	} else {
		return ips, cobra.ShellCompDirectiveNoFileComp
	}
}
