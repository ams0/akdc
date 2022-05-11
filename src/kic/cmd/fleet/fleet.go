// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"kic/boa"
	"kic/cmd/targets"

	"github.com/spf13/cobra"
)

var (
	FleetCmd = &cobra.Command{
		Use:   "flt",
		Short: "Retail Edge CLI",
		Long:  "Retail Edge CLI\n\n  A CLI for automating many Kubernetes fleet tasks",
	}

	// option variable
	grep string
)

func LoadCommands() *cobra.Command {
	if len(FleetCmd.Commands()) == 0 {
		FleetCmd.AddCommand(AzCmd)
		FleetCmd.AddCommand(CheckCmd)
		FleetCmd.AddCommand(CreateCmd)
		FleetCmd.AddCommand(DeleteCmd)
		FleetCmd.AddCommand(DnsCmd)
		FleetCmd.AddCommand(ExecCmd)
		FleetCmd.AddCommand(ListCmd)
		FleetCmd.AddCommand(NewAppCmd)
		FleetCmd.AddCommand(SshCmd)
		FleetCmd.AddCommand(targets.TargetsCmd)

		boa.AddScriptCommand(FleetCmd, "env", "List the environment variables", "env | grep AKDC | sort")

		FleetCmd.AddCommand(boa.AddFltCommand("curl", "curl the specified endpoint on each cluster", "", "curl"))
		FleetCmd.AddCommand(boa.AddFltCommand("arc-token", "Get Arc token from each cluster", "", "arc-token"))
		FleetCmd.AddCommand(boa.AddFltCommand("patch", "Run a patch command on each cluster", "", "patch"))
		FleetCmd.AddCommand(boa.AddFltCommand("pull", "Git pull the akdc repo", "", "pull"))
		FleetCmd.AddCommand(boa.AddFltCommand("sync", "Sync (reconcile) flux on each cluster", "", "sync"))
	}

	return FleetCmd
}
