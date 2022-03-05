// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package targets

import (
	"encoding/json"
	"fmt"
	"kic/boa"
	"kic/boa/cfmt"
	"os"

	"github.com/spf13/cobra"
)

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
	boa.ShellExecOut("git pull")

	content, err := os.ReadFile(AutoGitOpsConfigFile)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	txt := string(content)

	var result map[string]interface{}

	err = json.Unmarshal([]byte(txt), &result)

	if err != nil {
		cfmt.Error("unmarshal json faile")
		fmt.Println(err)
		return nil
	}

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
