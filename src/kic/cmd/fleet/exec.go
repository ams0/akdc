// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"fmt"
	"github.com/spf13/cobra"
	"kic/utils"
)

// execCmd runs a bash command on each server
var ExecCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a bash command on each server",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: akdc exec bashCommand")
			return
		}

		command := fmt.Sprintf("%s", args)

		if len(command) < 3 {
			fmt.Println("Usage: akdc exec bashCommand")
			return
		}

		command = command[1 : len(command)-1]

		utils.ExecClusters(command, grep)
	},
}

func init() {
	ExecCmd.Flags().StringVarP(&grep, "grep", "g", "", "grep conditional to filter by host")
}
