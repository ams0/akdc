// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package targets

import (
	"kic/boa"
	"kic/boa/cfmt"

	"github.com/spf13/cobra"
)

// PushCmd adds, commits, and pushes the GitOps targets to the repo
var PushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push the changes to GitHub",
	Args:  argsTargets,
	RunE:  runTargetsPushE,
}

func runTargetsPushE(cmd *cobra.Command, args []string) error {
	// make sure the repo is up to date
	if res, _ := boa.ShellExecOut("git status -s"); res == "" {
		cfmt.Info("nothing to push")
	} else {
		cfmt.Info("pulling from GitHub")
		boa.ShellExecE("git pull")

		cfmt.Info("committing changes to GitHub")
		boa.ShellExecE("git commit -am 'updated targets with akdc CLI'")

		cfmt.Info("pushing changes to GitHub")
		boa.ShellExecE("git push")
	}

	return nil
}
