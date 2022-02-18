// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var targetsPush bool

// checkCmd adds check subcommands
var targetsCmd = &cobra.Command{
	Use:   "targets",
	Short: "manage GitOps targets",
	Long:  `manage GitOps targets`,
}

func init() {
	targetsCmd.AddCommand(targetsAddCmd)
	targetsCmd.AddCommand(targetsClearCmd)
	targetsCmd.AddCommand(targetsListCmd)
	targetsCmd.AddCommand(targetsPushCmd)
	targetsCmd.AddCommand(targetsRemoveCmd)
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
	shellExecOut("git pull")

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
