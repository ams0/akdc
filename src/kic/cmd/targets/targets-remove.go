// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package targets

import (
	"fmt"
	"kic/boa/cfmt"

	"github.com/spf13/cobra"
)

// RemoveCmd removes a target from GitOps
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a GitOps target",
	Args:  argsTargets,
	RunE:  runTargetsRemoveE,
}

func runTargetsRemoveE(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return cfmt.ErrorE("Error: specify the target to remove")
	}

	result := getAutoGitOpsConfigMap()

	if result != nil {
		return cfmt.ErrorE("Error: unable to read autogitops.json")
	} else {
		t := result["targets"]

		if t == nil {
			fmt.Println("targets is empty")
			return nil
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

	return nil
}
