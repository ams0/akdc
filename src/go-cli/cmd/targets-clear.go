// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// checkFluxCmd checks each cluster for flux-check namespace
var targetsClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clear all GitOps targets",
	Long:  `clear all GitOps targets`,
	Args: func(cmd *cobra.Command, args []string) error {
		return checkForConfigFile()
	},
	Run: func(cmd *cobra.Command, args []string) {
		result := getAutoGitOpsConfigMap()

		if result != nil {
			var nt []interface{}

			result["targets"] = nt

			saveAutoGitOpsConfig(result)

			fmt.Println("targets cleared")
		}
	},
}
