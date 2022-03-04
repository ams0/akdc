// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"github.com/spf13/cobra"
	"kic/cmd/fleet/check"
)

// option variables
// common across several commands
var debug bool

// used by check, exec, sync and test commands
var grep string

// checkCmd adds check subcommands
var FleetCmd = &cobra.Command{
	Use:   "fleet",
	Short: "Commands for fleet of clusters",
}

func init() {
	FleetCmd.AddCommand(check.CheckCmd)
	FleetCmd.AddCommand(CreateCmd)
	FleetCmd.AddCommand(DeleteCmd)
	FleetCmd.AddCommand(ExecCmd)
	FleetCmd.AddCommand(GroupsCmd)
	FleetCmd.AddCommand(SshCmd)
	FleetCmd.AddCommand(SyncCmd)
	FleetCmd.AddCommand(TokenCmd)
	FleetCmd.AddCommand(PullCmd)
	FleetCmd.AddCommand(PatchCmd)
}
