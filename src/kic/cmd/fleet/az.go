// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"kic/boa"

	"github.com/spf13/cobra"
)

// ListCmd lists the clusters in the fleet
var AzCmd = &cobra.Command{
	Use:   "az",
	Short: "Azure CLI commands (customized)",
}

func init() {
	boa.AddCommandIfNotExist(AzCmd, boa.CreateScriptCommand("groups", "List Azure groups", fltAzGroupsScript()))
	boa.AddCommandIfNotExist(AzCmd, boa.CreateScriptCommand("login", "Login to Azure using the project's Service Principal", fltAzLoginScript()))
	boa.AddCommandIfNotExist(AzCmd, boa.CreateScriptCommand("logout", "Logout of Azure", fltAzLogoutScript()))
	boa.AddCommandIfNotExist(AzCmd, boa.CreateScriptCommand("vms", "List Azure VMs in the resource group", fltAzVMsScript()))
}

func fltAzGroupsScript() string {
	return "az group list -o table | sort | grep -e central- -e east- -e west- -e corp- -e retail-edge -e fleet"
}

func fltAzLoginScript() string {
	return "az login --service-principal --username $AKDC_SP_ID --tenant $AKDC_TENANT --password $AKDC_SP_KEY"
}

func fltAzLogoutScript() string {

	return "az logout"
}

func fltAzVMsScript() string {
	return `
		rg=$1

		if [ "$rg" = "" ]
		then
			rg="$FLT_RG"
		fi

		if [ "$rg" = "" ]
		then
			rg="$(git branch --show-current)"
		fi

		if [ "$rg" = "" ]
		then
			echo "usage: flt az vms resourceGroup"
			exit 0
		fi

		echo ""
		echo "getting VMs in resource group: $rg"
		echo ""

		hdrsort()
		{
			read -r
			printf "%s\\n" "$REPLY"
			sort
		}

		az vm list --query '[].name' -o table -g $rg | hdrsort
	`
}
