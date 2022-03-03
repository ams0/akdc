// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package targets

import (
	"encoding/json"
	"fmt"
	"kic/utils"
	"os"

	"github.com/spf13/cobra"
)

var targetsPush bool

// checkCmd adds check subcommands
var TargetsCmd = &cobra.Command{
	Use:   "targets",
	Short: "Manage GitOps targets",
}

func init() {
	TargetsCmd.AddCommand(AddCmd)
	TargetsCmd.AddCommand(ClearCmd)
	TargetsCmd.AddCommand(ListCmd)
	TargetsCmd.AddCommand(PushCmd)
	TargetsCmd.AddCommand(RemoveCmd)
}

var AutoGitOpsConfigFile = "./autogitops/autogitops.json"

func checkForConfigFile() error {
	if _, err := os.Stat(AutoGitOpsConfigFile); err != nil {
		return fmt.Errorf("GitOps file not found - please cd to an app with GitOps setup")
	}

	return nil
}

func getAutoGitOpsConfigMap() map[string]interface{} {
	// make sure the repo is up to date
	utils.ShellExecOut("git pull")

	content, err := os.ReadFile(AutoGitOpsConfigFile)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	txt := string(content)

	var result map[string]interface{}

	json.Unmarshal([]byte(txt), &result)

	return result
}

func saveAutoGitOpsConfig(result map[string]interface{}) {
	val, err := json.MarshalIndent(result, "", "    ")

	if err != nil {
		fmt.Println("Error saving:", err)
	} else {
		os.WriteFile(AutoGitOpsConfigFile, val, 0644)
	}
}
