// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// createCmd creates a new cluster
// todo - implement
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new cluster",
	Long:  `create a new cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("not implemented")
	},
}
