// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"fmt"
	"kic/boa"
	"kic/boa/cfmt"
	"strings"

	"github.com/spf13/cobra"
)

// ListCmd lists the clusters in the fleet
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List the clusters in the fleet",
	RunE:  runFleetList,
}

func init() {
	ListCmd.Flags().StringVarP(&grep, "grep", "g", "", "grep conditional to filter by host")
}

// run the kic fleet list command
func runFleetList(cmd *cobra.Command, args []string) error {
	cfmt.Info("Clusters in the fleet")
	fmt.Println()

	hostIPs, err := boa.ReadHostIPs("")

	if err != nil {
		return err
	}

	for _, line := range hostIPs {
		if len(line) < 30 {
			line = strings.Replace(line, "\t", "\t\t", -1)
		}
		fmt.Println(line)
	}

	return nil
}
