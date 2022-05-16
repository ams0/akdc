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

func LoadCommands(parent *cobra.Command) *cobra.Command {
	boa.AddCommandIfNotExist(parent, AzCmd)
	boa.AddCommandIfNotExist(parent, CheckCmd)
	boa.AddCommandIfNotExist(parent, CreateCmd)
	boa.AddCommandIfNotExist(parent, DeleteCmd)
	boa.AddCommandIfNotExist(parent, DnsCmd)
	boa.AddCommandIfNotExist(parent, ExecCmd)
	boa.AddCommandIfNotExist(parent, ListCmd)
	boa.AddCommandIfNotExist(parent, SshCmd)
	boa.AddCommandIfNotExist(parent, targets.TargetsCmd)

	boa.AddCommandIfNotExist(parent, boa.CreateScriptCommand("env", "List the environment variables", "env | grep AKDC | sort"))

	boa.AddCommandIfNotExist(parent, boa.CreateFltCommand("curl", "curl the specified endpoint on each cluster", "", "curl"))
	boa.AddCommandIfNotExist(parent, boa.CreateFltCommand("patch", "Run a patch command on each cluster", "", "patch"))
	boa.AddCommandIfNotExist(parent, boa.CreateFltCommand("pull", "Git pull the akdc repo", "", "pull"))
	boa.AddCommandIfNotExist(parent, boa.CreateFltCommand("sync", "Sync (reconcile) flux on each cluster", "", "sync"))

	if az := boa.GetCommandByUse(parent, "az"); az != nil {
		boa.AddCommandIfNotExist(az, boa.CreateFltCommand("arc-token", "Get Arc token from each cluster", "", "arc-token"))
	}

	// add tab completion to flt check app
	if chk := boa.GetCommandByUse(parent, "check"); chk != nil {
		if app := boa.GetCommandByUse(chk, "app"); app != nil {
			app.ValidArgsFunction = validArgsFleetCheckApp
		}
	}

	return parent
}
