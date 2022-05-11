// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"kic/boa"

	"github.com/spf13/cobra"
)

// check the clusters in the fleet
var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Check cluster status",
}

func init() {
	checkAppCmd := boa.AddFltCommand("app", "Check app status on each cluster", "", "check-app")
	checkAppCmd.ValidArgsFunction = validArgsFleetCheckApp

	CheckCmd.AddCommand(checkAppCmd)

	CheckCmd.AddCommand(boa.AddFltCommand("flux", "Check flux status on each cluster", "", "check-flux"))
	CheckCmd.AddCommand(boa.AddFltCommand("heartbeat", "Check https heartbeat on each cluster", "", "check-heartbeat"))
	CheckCmd.AddCommand(boa.AddFltCommand("logs", "Check the cloudinit logs on each cluster", "", "check-logs"))
	CheckCmd.AddCommand(boa.AddFltCommand("setup", "Check setup status on each cluster", "", "check-setup"))
}

// validate flt check app arg
func validArgsFleetCheckApp(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	apps, err := boa.ReadCompletionFile("flt-check-app-completion")

	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	// only one arg
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	// sugest from the completion file
	return apps, cobra.ShellCompDirectiveNoFileComp
}
