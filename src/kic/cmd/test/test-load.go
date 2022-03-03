// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package test

import (
	"fmt"
	"kic/cfmt"
	"kic/utils"
	"os"

	"github.com/spf13/cobra"
)

// test-load command
var LoadCmd = &cobra.Command{
	Use:   "load",
	Short: "Run a load test on each cluster",

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

		path := utils.GetBinDir() + "/.kic/commands/" + command

		// execute the file with "bash -c" if it exists
		if _, err := os.Stat(path); err == nil {
			utils.ShellExec(fmt.Sprintf("%s %s", path, params))
		} else {
			cfmt.Error(err)
		}
	},
}

// add command specific options
func init() {
	LoadCmd.Flags().StringVarP(&fileLoad, "file", "f", "benchmark.json", "Test file to use")
	LoadCmd.Flags().IntVarP(&duration, "duration", "", 30, "Test duration (seconds)")
	LoadCmd.Flags().IntVarP(&delayStart, "delay-start", "", 0, "Delay test start (seconds)")
}
