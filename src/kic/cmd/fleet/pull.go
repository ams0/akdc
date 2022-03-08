// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"kic/boa"

	"github.com/spf13/cobra"
)

// PullCmd runs git pull on the akdc repo on each cluster
var PullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Git pull the akdc repo",
	RunE: func(cmd *cobra.Command, args []string) error {
		return (boa.ExecClusters("git -C gitops pull", grep))
	},
}

func init() {
	PullCmd.Flags().StringVarP(&grep, "grep", "g", "", "grep conditional to filter by host")
}
