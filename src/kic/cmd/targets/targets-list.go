// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package targets

import (
	"fmt"
	"github.com/spf13/cobra"
	"kic/cfmt"
)

// checkFluxCmd checks each cluster for flux-check namespace
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List current targets",

	Args: func(cmd *cobra.Command, args []string) error {
		return checkForConfigFile()
	},
	Run: func(cmd *cobra.Command, args []string) {
		result := getAutoGitOpsConfigMap()

		if result != nil {
			if result["targets"] == nil {
				cfmt.Info("targets", "is empty")
			} else {
				cfmt.Info("Targets")
				fmt.Println(result["targets"])
			}
		}
	},
}
