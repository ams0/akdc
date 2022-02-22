// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// checkFluxCmd checks each cluster for flux-check namespace
var targetsRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove a GitOps target",
	Long:  `remove a GitOps target`,
	Args: func(cmd *cobra.Command, args []string) error {
		return checkForConfigFile()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: specify the target to remove")
			return
		}

		result := getAutoGitOpsConfigMap()

		if result != nil {
			t := result["targets"]

			if t == nil {
				fmt.Println("targets is empty")
				return
			}

			var nt []interface{}

			for _, v := range t.([]interface{}) {
				if v != args[0] {
					nt = append(nt, v)
				}
			}

			result["targets"] = nt

			saveAutoGitOpsConfig(result)

			fmt.Println(nt)
		}
	},
}
