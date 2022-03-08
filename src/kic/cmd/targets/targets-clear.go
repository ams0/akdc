// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package targets

import (
	"fmt"
	"kic/boa/cfmt"

	"github.com/spf13/cobra"
)

// ClearCmd clears the GitOps targets
var ClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all GitOps targets",
	Args:  argsTargets,
	RunE:  runTargetsClearE,
}

func runTargetsClearE(cmd *cobra.Command, args []string) error {
	result := getAutoGitOpsConfigMap()

	if result != nil {
		var nt []interface{}

		result["targets"] = nt

		saveAutoGitOpsConfig(result)

		fmt.Println("targets cleared")

		return nil
	}

	return cfmt.ErrorE("Unable to read targets")
}
