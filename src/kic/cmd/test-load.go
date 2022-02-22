// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"kic/cfmt"
)

// test-load command
var testLoadCmd = &cobra.Command{
	Use:   "load",
	Short: "run a load test on each cluster",
	Long:  `run a load test on each cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		command := fmt.Sprintf("webv -r -s http://localhost:30080")

		cfmt.Info("Running load test:")
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
	},
}

// add command specific options
func init() {
	testLoadCmd.Flags().StringVarP(&fileLoad, "file", "f", "../webv/benchmark.json", "Test file to use")
	testLoadCmd.Flags().IntVarP(&duration, "duration", "", 30, "Test duration (seconds)")
	testLoadCmd.Flags().IntVarP(&delayStart, "delay-start", "", 0, "Delay test start (seconds)")
	testLoadCmd.Flags().IntVarP(&sleep, "sleep", "l", 100, "Sleep (ms) between each request")
}
