// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package utils

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

	"github.com/spf13/cobra"
)

// execute a bash command
func ShellExec(cmd string) {
	shell := exec.Command("bash", "-c", cmd)

	shell.Stdin = os.Stdin
	shell.Stdout = os.Stdout
	shell.Stderr = os.Stderr

	shell.Run()
}

// execute a bash command with args
func ShellExecArgs(cmd string, args []string) {
	command := cmd
	for _, arg := range args {
		command += " " + arg
	}
	fmt.Println(command)
	ShellExec(command)
}

// execute a bash command and return stdout
func ShellExecOut(cmd string) string {
	shell := exec.Command("bash", "-c", cmd)

	var out bytes.Buffer
	shell.Stdout = &out

	shell.Run()

	return out.String()
}

func ReadConfigValue(key string) string {
	key = strings.TrimSpace(key)

	if !strings.HasSuffix(key, ":") {
		key += ":"
	}

	config := os.Getenv("HOME") + "/.kic"

	if _, err := os.Stat(config); err == nil {
		lines := ReadLinesFromFile(config)

		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, key) {
				line = strings.TrimSpace(strings.Replace(line, key, "", 1))
				return line
			}
		}
	}

	return ""
}

// read the ips file
func ReadHostIPs(grep string) []string {
	command := ""

	if _, err := os.Stat("./ips"); err != nil {
		file := ReadConfigValue("defaultIPs:")
		if file != "" {
			command = "cat " + file + " | sort"
		}
	} else {
		command = "cat ips | sort"
	}

	if command == "" {
		cfmt.ExitErrorMessage("fleet file not found")
	}

	if grep != "" {
		err := CheckForBadChars(grep, "grep")
		if err != nil {
			cfmt.ExitErrorMessage(err)
		}

		command += " | grep " + grep
	}

	lines := strings.Split(string(ShellExecOut(command)), "\n")

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
func ExecClusters(cmd string, grep string) {
	hostIPs := ReadHostIPs(grep)

	ch := make(chan string)

	for _, hostIP := range hostIPs {
		cols := strings.Split(hostIP, "\t")

		if len(cols) > 1 {
			go ExecCluster(cols[0], cols[1], cmd, ch)
		}
	}

	// todo - add timeout
	for i := 0; i < len(hostIPs); i++ {
		<-ch
	}
}

// run a command on one cluster via ssh
func ExecCluster(host string, ip string, cmd string, ch chan string) {
	cmd = fmt.Sprintf("ssh -p 2222 -o \"StrictHostKeyChecking=no\" -o ConnectTimeout=5 akdc@%s %s", ip, cmd)

	ShellExec(cmd)

	ch <- host
}

// check for dangerous characters sent to bash
func CheckForBadChars(source string, param string) error {

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

// get the path to the executable's directory
func GetBinDir() string {
	ex, err := os.Executable()

	if err != nil {
		log.Fatal(err)
	}

	// get the parent of bin
	return filepath.Dir(ex)
}

// get the file name from the executing directory
func GetBinName() string {
	ex, err := os.Executable()

	if err != nil {
		log.Fatal(err)
	}

	// get the parent of bin
	return filepath.Base(ex)
}

// get the path to the modules (i.e. /bin/kic/.kic)
func GetModPath() string {
	modPath := GetBinDir()
	app := GetBinName()

	if strings.HasPrefix(app, "__debug") {
		// running in debugger - assume package name == source directory
		app = filepath.Base(modPath)
	}

	// complete the paths
	return modPath + "/." + app + "/"

}

// get the path to the repo base
func GetRepoBase() string {
	base := os.Getenv("REPO_BASE")

	if base == "" {
		ex, err := os.Executable()

		if err != nil {
			log.Fatal(err)
		}

		base = filepath.Dir(ex)
		base = filepath.Dir(base)

		if strings.HasSuffix(base, "src") {
			base = filepath.Dir(base)
		}
	}

	return base
}

func ReadTextFileFromBin(name string) string {
	dir := GetModPath()
	path := filepath.Join(dir, name)
	return ReadTextFile(path)
}

// read a file and return the text
func ReadTextFile(path string) string {
	content, err := os.ReadFile(path)

	if err != nil {
		return ""
	}

	return string(content)
}

// read lines from a text file
func ReadLinesFromFile(path string) []string {
	return strings.Split(ReadTextFile(path), "\n")
}

// execute a command in bin/.kic/commands
func ExecCommand(cmd string) {
	path := GetBinDir() + "/.kic/commands/" + cmd

	// execute the file with "bash -c" if it exists
	if _, err := os.Stat(path); err == nil {
		cfmt.Info("Running command: " + cmd)
		ShellExec(fmt.Sprintf("%s %s", path, os.Args))
	}
}

// execute a bash command
func ShellExecE(cmd string) error {
	shell := exec.Command("bash", "-c", cmd)

	shell.Stdin = os.Stdin
	shell.Stdout = os.Stdout
	shell.Stderr = os.Stderr

	return shell.Run()
}

// create a command that runs the bash command
func AddRunCommand(use string, short string, long string, command string) *cobra.Command {
	modPath := GetBinDir()
	app := GetBinName()

	if strings.HasPrefix(app, "__debug") {
		// running in debugger - assume package name == source directory
		app = filepath.Base(modPath)
	}

	// complete the paths
	modPath += "/." + app + "/"
	cmdPath := modPath + "commands/"

	runCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				ShellExecArgs(cmdPath+command, args)
			} else {
				ShellExec(cmdPath + command)
			}
		},
	}

	return runCmd
}
