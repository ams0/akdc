// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"kic/cfmt"
	"os"
)

// test-load command
var testLoadCmd = &cobra.Command{
	Use:   "load",
	Short: "run a load test on each cluster",

	Run: func(cmd *cobra.Command, args []string) {
		command := "test"

		cfmt.Info("Running load test")

		if sleep < 1 {
			sleep = 100
		}

		params := "--run-loop "
		params += getTestFlagValues()

		// add test-load specific options to command line
		if fileLoad != "" {
			params += " --files " + fileLoad
		}

		if random {
			params += " --random"
		}

		if duration > 0 {
			params += fmt.Sprintf(" --duration %d", duration)
		}

		if delayStart > 0 {
			params += fmt.Sprintf(" --delay-start %d", delayStart)
		}

		path := getParentDir() + "/.kic/commands/" + command

		// execute the file with "bash -c" if it exists
		if _, err := os.Stat(path); err == nil {
			shellExec(fmt.Sprintf("%s %s", path, params))
		} else {
			cfmt.Error(err)
		}
	},
}

// add command specific options
func init() {
	testLoadCmd.Flags().StringVarP(&fileLoad, "file", "f", "benchmark.json", "Test file to use")
	testLoadCmd.Flags().IntVarP(&duration, "duration", "", 30, "Test duration (seconds)")
	testLoadCmd.Flags().IntVarP(&delayStart, "delay-start", "", 0, "Delay test start (seconds)")
}
