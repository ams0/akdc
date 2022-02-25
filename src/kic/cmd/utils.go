// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"kic/cfmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
func readHostIPs(grep string) []string {
	command := "cat ips | sort"

	if grep != "" {
		err := checkForBadChars(grep, "grep")
		if err != nil {
			log.Fatal(err)
		}

		command += " | grep " + grep
	}

	lines := strings.Split(string(shellExecOut(command)), "\n")

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
func execClusters(cmd string, grep string) {
	hostIPs := readHostIPs(grep)

	ch := make(chan string)

	for _, hostIP := range hostIPs {
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

// check for dangerous characters sent to bash
func checkForBadChars(source string, param string) error {

	if source != "" {
		badChars := "|&;<>"

		for _, ch := range badChars {
			if strings.Contains(source, string(ch)) {
				return errors.New(fmt.Sprintf("Invalid character in parameter %s", param))
			}
		}
	}

	return nil
}

// get the path to the executable's parent
func getParentDir() string {
	ex, err := os.Executable()

	if err != nil {
		log.Fatal(err)
	}

	// get the parent of bin
	return filepath.Dir(ex)
}

// read a file and return the text
func readTextFile(path string) string {
	content, err := os.ReadFile(path)

	if err != nil {
		return ""
	}

	return string(content)
}

// execute a command in bin/.kic/commands
func execCommand(cmd string) {
	path := getParentDir() + "/.kic/commands/" + cmd

	// execute the file with "bash -c" if it exists
	if _, err := os.Stat(path); err == nil {
		cfmt.Info("Running command: " + cmd)
		shellExec(fmt.Sprintf("%s %s", path, os.Args))
	}
}
