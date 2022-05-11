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
	boa.AddScriptCommand(AzCmd, "groups", "List Azure groups", fltAzGroupsScript())
	boa.AddScriptCommand(AzCmd, "login", "Login to Azure using the project's Service Principal", fltAzLoginScript())
	boa.AddScriptCommand(AzCmd, "logout", "Logout of Azure", fltAzLogoutScript())
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
