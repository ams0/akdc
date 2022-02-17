// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// execCmd runs a bash command on each server
// todo - need to work on formatting for more complex commands
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "execute a bash command on each server",
	Long:  `execute a bash command on each server`,
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

		execClusters(command, grep)
	},
}

func init() {
	execCmd.Flags().StringVarP(&grep, "grep", "g", "", "grep conditional to filter by host")
}
