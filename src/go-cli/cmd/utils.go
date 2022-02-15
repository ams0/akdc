// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// execute a bash command
func shellExec(cmd string) {
	shell := exec.Command("bash", "-c", cmd)

	shell.Stdin = os.Stdin
	shell.Stdout = os.Stdout
	shell.Stderr = os.Stderr

	shell.Run()
}

// execute a bash command and return stdout
func shellExecOut(cmd string) string {
	shell := exec.Command("bash", "-c", cmd)

	var out bytes.Buffer
	shell.Stdout = &out

	shell.Run()

	return out.String()
}

// read the ips file
func readhostIPs() []string {
	content, err := ioutil.ReadFile("ips")

	if err != nil {
		return nil
	}

	lines := strings.Split(string(content), "\n")
	var ips []string = nil

	for _, line := range lines {
		ip := strings.Split(line, "\t")

		if len(ip) > 1 {
			ips = append(ips, line)
		}
	}

	return ips
}

// run a command on all clusters
func execClusters(cmd string) {
	hostIPs := readhostIPs()

	ch := make(chan string)

	for _, hostIP := range hostIPs {
		// todo - generalize and pass in function
		cols := strings.Split(hostIP, "\t")

		if len(cols) > 1 {
			go execCluster(cols[0], cols[1], cmd, ch)
		}
	}

	// todo - add timeout
	for i := 0; i < len(hostIPs); i++ {
		<-ch
	}
}

// run a command on one cluster via ssh
func execCluster(host string, ip string, cmd string, ch chan string) {
	cmd = fmt.Sprintf("ssh -p 2222 -o \"StrictHostKeyChecking=no\" -o ConnectTimeout=5 akdc@%s %s", ip, cmd)

	shellExec(cmd)

	ch <- host
}
