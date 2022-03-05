// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package boa

import (
	"bytes"
	"os"
	"os/exec"
)

// execute a bash command with args
func ShellExecArgsE(cmd string, args []string) error {
	command := cmd
	for _, arg := range args {
		command += " " + arg
	}
	return ShellExecE(command)
}

// execute a bash command and return stdout
func ShellExecOut(cmd string) (string, error) {
	shell := exec.Command("bash", "-c", cmd)

	var out bytes.Buffer
	shell.Stdout = &out

	err := shell.Run()

	return out.String(), err
}

// execute a bash command
func ShellExecE(cmd string) error {
	shell := exec.Command("bash", "-c", cmd)

	shell.Stdin = os.Stdin
	shell.Stdout = os.Stdout
	shell.Stderr = os.Stderr

	return shell.Run()
}
