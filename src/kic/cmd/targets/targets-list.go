// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package targets

import (
	"fmt"
	"kic/boa/cfmt"

	"github.com/spf13/cobra"
)

// ListCmd lists the GitOps targets
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List current targets",
	Args:  argsTargets,
	RunE:  runTargetsListE,
}

func runTargetsListE(cmd *cobra.Command, args []string) error {
	result := getAutoGitOpsConfigMap()

	if result != nil {
		if result["targets"] == nil {
			cfmt.Info("targets is empty")
		} else {
			cfmt.Info("Targets")
			fmt.Println(result["targets"])
		}
		return nil
	} else {
		return cfmt.ErrorE("failed to read autogitops.json")
	}
}
