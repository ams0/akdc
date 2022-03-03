// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package targets

import (
	"fmt"
	"github.com/spf13/cobra"
)

// checkFluxCmd checks each cluster for flux-check namespace
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a GitOps target",

	Args: func(cmd *cobra.Command, args []string) error {
		return checkForConfigFile()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: specify the target to add")
			return
		}

		result := getAutoGitOpsConfigMap()

		if result != nil {
			t := result["targets"]

			var nt []interface{}

			found := false

			if t != nil {
				for _, v := range t.([]interface{}) {
					nt = append(nt, v)

					if v == args[0] {
						found = true
					}
				}
			}

			if !found {
				nt = append(nt, args[0])
			}

			result["targets"] = nt

			saveAutoGitOpsConfig(result)

			fmt.Println(nt)
		}
	},
}
