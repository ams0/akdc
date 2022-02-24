// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get information from the cluster",
	Long:  `get information from the cluster`,
}

func init() {
	getCmd.AddCommand(getTokenCmd)
}
