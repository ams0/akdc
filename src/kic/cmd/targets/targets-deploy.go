// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package targets

import (
	"kic/boa"
	"kic/boa/cfmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// DeployCmd adds, commits, and deploys the GitOps targets to the repo
var DeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy to the target stores",
	Args:  argsTargets,
	RunE:  runTargetsDeployE,
}

// run targets deploy cmd
func runTargetsDeployE(cmd *cobra.Command, args []string) error {
	// build the commit message from app name
	app, _ := os.Getwd()
	app = filepath.Base(app)

	// make sure repo is up-to-date
	if err := boa.ShellExecE("git pull"); err != nil {
		cfmt.ErrorE("git pull failed")
		return nil
	}

	// make sure the repo is up to date
	res, err := boa.ShellExecOut("git status -s", false)

	if err != nil {
		return cfmt.ErrorE(err)
	}

	if res != "" {
		// commit and push
		commit := "git commit -am 'Secure Build: " + app + "'"
		if err := boa.ShellExecE(commit); err != nil {
			cfmt.ErrorE("git commit failed")
			return nil
		}
	}

	if err := boa.ShellExecE("git push"); err != nil {
		cfmt.ErrorE("git push failed")
		return nil
	}

	cfmt.Info("updated " + app)

	return nil
}
