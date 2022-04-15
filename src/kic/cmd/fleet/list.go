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
		cols := strings.Split(line, "\t")

		if len(cols) > 1 {
			fmt.Print(cols[0])
			if len(cols[0]) < 30 {
				fmt.Print(strings.Repeat(" ", 30-len(cols[0])))
			}
			fmt.Println(cols[1])
		} else {
			fmt.Println(line)
		}
	}

	return nil
}
