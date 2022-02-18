// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

// test-load command
var testLoadCmd = &cobra.Command{
	Use:   "load",
	Short: "run a load test on each cluster",
	Long:  `run a load test on each cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		hostIPs := readHostIPs(grep)

		command := fmt.Sprintf("webv -r -s")
		servers := ""

		for _, line := range hostIPs {
			cols := strings.Split(line, "\t")
			servers += fmt.Sprintf(" https://%s.%s", cols[0], "cseretail.com")
		}

		if len(servers) > 1 {
			fmt.Println("\nRunning load test:\n")
			command += servers

			command += getTestFlagValues()

			// add test-load specific options to command line
			if fileLoad != "" {
				command += " --files " + fileLoad
			}

			if random {
				command += " --random"
			}

			if duration > 0 {
				command += fmt.Sprintf(" --duration %d", duration)
			}

			if delayStart > 0 {
				command += fmt.Sprintf(" --delay-start %d", delayStart)
			}

			shellExec(command)
		}
	},
}

// add command specific options
func init() {
	testLoadCmd.Flags().StringVarP(&fileLoad, "file", "f", "./webv/benchmark.json", "Test file to use")
	testLoadCmd.Flags().IntVarP(&duration, "duration", "", 10, "Test duration (seconds)")
	testLoadCmd.Flags().IntVarP(&delayStart, "delay-start", "", 0, "Delay test start (seconds)")
	testLoadCmd.Flags().IntVarP(&sleep, "sleep", "l", 0, "Sleep (ms) between each request")
}
