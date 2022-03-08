// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"fmt"
	"kic/boa"
	"kic/boa/cfmt"

	"github.com/spf13/cobra"
)

// execCmd runs a bash command on each server
var ExecCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a bash command on each server",
	Args:  cobra.MinimumNArgs(1),
	RunE:  runFleetExecE,
}

func init() {
	ExecCmd.Flags().StringVarP(&grep, "grep", "g", "", "grep conditional to filter by host")
}

// run kic fleet exec command
func runFleetExecE(cmd *cobra.Command, args []string) error {
	command := fmt.Sprintf("%s", args)

	// command will have []
	if len(command) < 3 {
		return cfmt.ErrorE("Usage: flt exec bashCommand")
	}

	// remove []
	command = command[1 : len(command)-1]

	return (boa.ExecClusters(command, grep))
}
