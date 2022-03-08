// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package targets

import (
	"fmt"
	"kic/boa/cfmt"

	"github.com/spf13/cobra"
)

// AddCmd adds a target to GitOps
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a GitOps target",

	Args: argsTargets,

	RunE: runTargetsAddE,
}

func runTargetsAddE(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return cfmt.ErrorE("Error: specify the target to add")
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

	return nil
}
