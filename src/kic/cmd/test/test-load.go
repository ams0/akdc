// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package test

import (
	"fmt"
	"kic/boa"
	"kic/boa/cfmt"
	"os"

	"github.com/spf13/cobra"
)

// command line options
var fileLoad string
var delayStart int
var duration int
var random bool

// test-load command
var LoadCmd = &cobra.Command{
	Use:   "load",
	Short: "Run a load test on each cluster",

	Run: func(cmd *cobra.Command, args []string) {
		cfmt.Info("Running load test")

		if sleep < 1 {
			sleep = 100
		}

		// get shared options
		params := getTestFlagValues()

		// add test-load specific options to command line
		params += " --run-loop "

		if fileLoad != "" {
			params += " --files " + fileLoad
		}

		if delayStart > 0 {
			params += fmt.Sprintf(" --delay-start %d", delayStart)
		}

		if duration > 0 {
			params += fmt.Sprintf(" --duration %d", duration)
		}

		if random {
			params += " --random"
		}

		// build the path to the script
		path := boa.GetBinDir() + "/.kic/commands/test"

		// execute the file with "bash -c" if it exists
		if _, err := os.Stat(path); err == nil {
			boa.ShellExecE(fmt.Sprintf("%s %s", path, params))
		} else {
			cfmt.Error(err)
		}
	},
}

// add command specific options
func init() {
	LoadCmd.Flags().StringVarP(&fileLoad, "file", "f", "benchmark.json", "Test file to use")
	LoadCmd.Flags().IntVarP(&delayStart, "delay-start", "", 0, "Delay test start (seconds)")
	LoadCmd.Flags().IntVarP(&duration, "duration", "", 30, "Test duration (seconds)")
	LoadCmd.Flags().BoolVarP(&random, "random", "", false, "Randomize tests")
}
