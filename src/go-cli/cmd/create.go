// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path"
	"strings"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new cluster",
	Long:  `create a new cluster`,

	Args: func(cmd *cobra.Command, args []string) error {
		pat := os.Getenv("HOME") + "/.ssh/akdc.pat"

		if _, err := os.Stat(pat); err != nil {
			return fmt.Errorf("Please export your GitOps PAT to %s before running akdc create", pat)
		}

		if ssl && zone == "" {
			return fmt.Errorf("you must specify --zone to use --ssl")
		}

		if ssl && pem == "" {
			return fmt.Errorf("you must specify --pem to use --ssl")
		}

		if ssl && key == "" {
			return fmt.Errorf("you must specify --key to use --ssl")
		}

		if group == "" {
			group = cluster
		}

		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		createGroup()

		ip := createVM()

		if ip != "" {
			createDNS(ip)
		}

		// remove the cluster template
		os.Remove("cluster-" + cluster + ".sh")

		if ip != "" {
			fmt.Println("\nCluster created")
		}
	},
}

// add command specific flags
func init() {
	createCmd.Flags().StringVarP(&cluster, "cluster", "c", "", "Kubernetes cluster name (required)")
	createCmd.MarkFlagRequired("cluster")
	createCmd.Flags().StringVarP(&group, "group", "g", "", "Azure resource group name")
	createCmd.Flags().StringVarP(&location, "location", "l", "centralus", "Azure location")
	createCmd.Flags().StringVarP(&repo, "repo", "r", "retaildevcrews/edge-gitops", "GitOps repo name")
	createCmd.Flags().StringVarP(&zone, "zone", "z", "", "DNS domain name")
	createCmd.Flags().BoolVarP(&ssl, "ssl", "s", false, "Use SSL cert (must specify --zone)")
	createCmd.Flags().StringVarP(&pem, "pem", "p", "~/.ssh/certs.pem", "Path to SSL .pem file")
	createCmd.Flags().StringVarP(&key, "key", "k", "~/.ssh/certs.key", " Path to SSL .key file")
	createCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Quiet mode")
}

// get the path to repo/src/cli
func getTemplatePath() string {
	ex, err := os.Executable()

	if err != nil {
		log.Fatal(err)
	}

	dir := path.Dir(ex)

	return dir + "/akdc.templ"
}

// create Azure Resource Group
func createGroup() {
	fmt.Println("Creating Azure Resource Group")

	os.Remove("cluster-" + cluster + ".sh")

	command := "sed \"s/{{cluster}}/" + cluster + "/g\" " + getTemplatePath() + " | "
	command += "sed \"s/{{fqdn}}/" + cluster + "." + zone + "/g\" | "
	command += "sed \"s~{{repo}}~" + repo + "~g\" "
	command += "> cluster-" + cluster + ".sh"
	shellExec(command)

	rgTags := "akdc=true server=" + cluster

	if zone != "" {
		rgTags += " zone=" + zone
	}

	command = "az group create -l " + location + " -n " + group + " -o table --tags " + rgTags
	shellExec(command)
}

// create Azure VM
func createVM() string {
	fmt.Println("Creating Azure VM")

	command := "az vm create \\\n"
	command += " -g " + group + " \\\n"
	command += " -l " + location + " \\\n"
	command += " -n " + cluster + " \\\n"
	command += " --admin-username akdc \\\n"
	command += " --size standard_D2as_v5 \\\n"
	command += " --image Canonical:0001-com-ubuntu-server-focal:20_04-lts-gen2:latest \\\n"
	command += " --os-disk-size-gb 128 \\\n"
	command += " --storage-sku Premium_LRS \\\n"
	command += " --generate-ssh-keys \\\n"
	command += " --public-ip-sku Standard \\\n"
	command += " --query publicIpAddress \\\n"
	command += " -o tsv \\\n"
	command += " --custom-data cluster-" + cluster + ".sh"

	ip := strings.TrimSpace(shellExecOut(command))

	if ip == "" {
		fmt.Println("Failed to create cluster")
		return ""
	}
	fmt.Println("Cluster created:", cluster, ip)

	f, err := os.OpenFile("ips", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}

	if _, err := f.WriteString(fmt.Sprintf("%s\t%s\n", cluster, ip)); err != nil {
		log.Println(err)
	}

	f.Close()

	fmt.Println("\nDeleting default nsg\n")
	command = "az network nsg rule delete -g " + group + " --nsg-name " + cluster + "NSG -o table --name default-allow-ssh"
	shellExec(command)

	fmt.Println("\nCreating SSH rule on port 2222\n")

	command = "az network nsg rule create \\\n"
	command += "-g " + group + " \\\n"
	command += "--nsg-name " + cluster + "NSG \\\n"
	command += "-n SSH-http \\\n"
	command += "--description \"SSH http https\" \\\n"
	command += "--destination-port-ranges 2222 80 443 \\\n"
	command += "--protocol tcp \\\n"
	command += "--access allow \\\n"
	command += "--priority 1202 -o table"
	shellExec(command)

	return ip
}

// create DNS entry and copy SSL certs
func createDNS(ip string) {
	if zone != "" {
		fmt.Println("Creating DNS Entry")

		command := "az network dns record-set a list \\\n"
		command += "--query \"[?name=='" + cluster + "'].{IP:aRecords}\" \\\n"
		command += "--resource-group tld \\\n"
		command += "--zone-name " + zone + " -o json | jq -r '.[].IP[].ipv4Address'"

		oldIP := strings.TrimSpace(shellExecOut(command))

		command = "az network dns record-set a add-record \\\n"
		command += "-g tld \\\n"
		command += "-z " + zone + " \\\n"
		command += "-n " + cluster + " \\\n"
		command += "-a " + ip + " \\\n"
		command += "--ttl 10 -o table"
		shellExec(command)

		if oldIP != "" && oldIP != ip {
			fmt.Println("Removing old IP:", ip)

			command = "az network dns record-set a remove-record \\\n"
			command += "-g tld \\\n"
			command += "-z " + zone + " \\\n"
			command += "-n " + cluster + " \\\n"
			command += "-a " + oldIP + " -o table"
			shellExec(command)
		}

		if ssl {
			fmt.Println("\nCopying SSL certs\n")

			shellExec("scp -P 2222 -o \"StrictHostKeyChecking=no\" -o ConnectTimeout=600 ~/.ssh/certs.pem akdc@" + ip + ":~/.ssh/certs.pem")
			shellExec("scp -P 2222 -o \"StrictHostKeyChecking=no\" -o ConnectTimeout=60 ~/.ssh/certs.key akdc@" + ip + ":~/.ssh/certs.key")
			shellExec("scp -P 2222 -o \"StrictHostKeyChecking=no\" -o ConnectTimeout=60 ~/.ssh/akdc.pat akdc@" + ip + ":~/.ssh/akdc.pat")
		}
	}
}
