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
	// make sure the repo is up to date
	if res, _ := boa.ShellExecOut("git status -s"); res == "" {
		cfmt.Info("no changes found")
	} else {
		// build the commit message from app name
		app, _ := os.Getwd()
		app = filepath.Base(app)
		commit := "git commit -am 'Secure Build: " + app + "'"

		// make sure repo is up-to-date
		boa.ShellExecOut("git pull")

		// commit and push
		boa.ShellExecOut(commit)
		boa.ShellExecOut("git push")

		cfmt.Info("updated " + app)
	}

	return nil
}
