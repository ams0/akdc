// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"akdc/cfmt"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new cluster",
	Long:  `create a new cluster`,

	Args: func(cmd *cobra.Command, args []string) error {
		pat := os.Getenv("HOME") + "/.ssh/akdc.pat"

		// make sure personal access token exists
		if _, err := os.Stat(pat); err != nil {
			return fmt.Errorf("Please export your GitOps PAT to %s before running akdc create", pat)
		}

		// validate ssl and domain
		if ssl != "" {
			if pem == "" {
				return fmt.Errorf("you must specify --pem to use --ssl")
			}

			if key == "" {
				return fmt.Errorf("you must specify --key to use --ssl")
			}

			if len(ssl) < 3 {
				return fmt.Errorf("invalid --ssl")
			}

			if !strings.Contains(ssl, ".") {
				return fmt.Errorf("invalid --ssl")
			}
		}

		// default resource group is cluster name
		if group == "" {
			group = cluster
		}

		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		// create the azure resource group
		createGroup()

		// create the vm and get the IP
		ip := createVM()

		// fail if createVM fails
		if ip != "" {
			// create DNS entry
			createDNS(ip)

			// copy files to VM
			scpFilesToVM(ip, ssl)
		}

		// remove the cluster template
		os.Remove("cluster-" + cluster + ".sh")

		// success
		if ip != "" {
			cfmt.Info("VM Configured")
		}
	},
}

// check to see if the Azure Resource Group exists
func groupExists() bool {
	ex := shellExecOut("az group exists -g " + group)
	return strings.HasPrefix(ex, "true")
}

// check to see if the VM exists in the RG
func vmExists() bool {
	command := fmt.Sprintf("az vm show -g %s --name %s --query 'name' -o tsv", group, cluster)
	res := strings.TrimSpace(shellExecOut(command))
	return strings.ToLower(cluster) == strings.ToLower(res)
}

var dapr bool
var arcEnabled bool

// add akdc create specific flags
func init() {
	createCmd.Flags().StringVarP(&cluster, "cluster", "c", "", "Kubernetes cluster name (required)")
	createCmd.MarkFlagRequired("cluster")
	createCmd.Flags().StringVarP(&group, "group", "g", "", "Azure resource group name")
	createCmd.Flags().StringVarP(&location, "location", "l", "centralus", "Azure location")
	createCmd.Flags().StringVarP(&repo, "repo", "r", "retaildevcrews/edge-gitops", "GitOps repo name")
	createCmd.Flags().StringVarP(&ssl, "ssl", "s", "", "SSL domain name")
	createCmd.Flags().StringVarP(&pem, "pem", "p", "~/.ssh/certs.pem", "Path to SSL .pem file")
	createCmd.Flags().StringVarP(&key, "key", "k", "~/.ssh/certs.key", " Path to SSL .key file")
	createCmd.Flags().BoolVarP(&dapr, "dapr", "", false, "Install Dapr and Radius")
	createCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Quiet mode")
	createCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Create VM in debug mode")
	createCmd.Flags().BoolVarP(&arcEnabled, "arc", "a", false, "Connect kubernetes cluster to Azure via Azure ARC")
}

// get the path to the executable's parent
func getParentDir() string {
	ex, err := os.Executable()

	if err != nil {
		log.Fatal(err)
	}

	// get the parent of bin
	return filepath.Dir(filepath.Dir(ex))
}

// get the path to template file
func getTemplatePath() string {
	return getParentDir() + "/vm-setup-files/akdc.templ"
}

// create Azure Resource Group
func createGroup() {
	// fail if the Azure VM exists
	if vmExists() {
		cfmt.Error("Azure VM exists")
		fmt.Println("\nPlease use a different VM or delete the VM")
		os.Exit(2)
	}

	// don't create the group if it exists
	if groupExists() {
		return
	}

	cfmt.Info("Creating Azure Resource Group")

	rgTags := "akdc=true server=" + cluster

	if ssl != "" {
		rgTags += " zone=" + ssl
	}

	command := "az group create -l " + location + " -n " + group + " -o table --tags " + rgTags

	err := shellExecE(command)
	if err != nil {
		cfmt.ExitError(err)
	}
}

// create vm setup script
func createVMSetupScript() {
	os.Remove("cluster-" + cluster + ".sh")

	// create the custom VM script
	content, err := os.ReadFile(getTemplatePath())

	if err != nil {
		cfmt.Error(err)
		return
	}

	// create the vm setup script from the template
	command := string(content)
	command = strings.Replace(command, "{{cluster}}", cluster, -1)
	command = strings.Replace(command, "{{dapr}}", strconv.FormatBool(dapr), -1)
	command = strings.Replace(command, "{{debug}}", strconv.FormatBool(debug), -1)
	command = strings.Replace(command, "{{fqdn}}", cluster+"."+ssl, -1)
	command = strings.Replace(command, "{{repo}}", repo, -1)
	command = strings.Replace(command, "{{group}}", group, -1)
	command = strings.Replace(command, "{{arcEnabled}}", strconv.FormatBool(arcEnabled), -1)
	os.WriteFile("cluster-"+cluster+".sh", []byte(command), 0644)
}

// create Azure VM
func createVM() string {
	cfmt.Info("Creating Azure VM")
	managedIdentityID := os.Getenv("AKDC_MI")

	// create the setup script from the template
	createVMSetupScript()

	command := "az vm create \\\n"
	command += " -g " + group + " \\\n"
	command += " -l " + location + " \\\n"
	command += " -n " + cluster + " \\\n"
	command += " --admin-username akdc \\\n"
	if managedIdentityID != "" {
		cfmt.Info("Assigning Managed Identity to VM")
		command += " --assign-identity " + managedIdentityID + "\\\n"
	}
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
		cfmt.Error("Failed to create cluster")
		return ""
	}
	cfmt.Info("VM created")
	fmt.Println(cluster, ip)

	f, err := os.OpenFile("ips", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		cfmt.Error(err)
	}

	if _, err := f.WriteString(fmt.Sprintf("%s\t%s\n", cluster, ip)); err != nil {
		log.Println(err)
	}

	f.Close()

	cfmt.Info("Deleting default nsg")
	command = "az network nsg rule delete -g " + group + " --nsg-name " + cluster + "NSG -o table --name default-allow-ssh"
	shellExecOut(command)

	cfmt.Info("Creating SSH rule on port 2222")

	command = "az network nsg rule create \\\n"
	command += "-g " + group + " \\\n"
	command += "--nsg-name " + cluster + "NSG \\\n"
	command += "-n SSH-http \\\n"
	command += "--description \"SSH http https\" \\\n"
	command += "--destination-port-ranges 2222 80 443 \\\n"
	command += "--protocol tcp \\\n"
	command += "--access allow \\\n"
	command += "--priority 1202 -o table"
	shellExecOut(command)

	return ip
}

// create DNS entry and copy SSL certs
func createDNS(ip string) {
	if ssl != "" {
		cfmt.Info("Creating DNS Entry")

		command := "az network dns record-set a list \\\n"
		command += "--query \"[?name=='" + cluster + "'].{IP:aRecords}\" \\\n"
		command += "--resource-group tld \\\n"
		command += "--zone-name " + ssl + " -o json | jq -r '.[].IP[].ipv4Address'"

		oldIP := strings.TrimSpace(shellExecOut(command))

		command = "az network dns record-set a add-record \\\n"
		command += "-g tld \\\n"
		command += "-z " + ssl + " \\\n"
		command += "-n " + cluster + " \\\n"
		command += "-a " + ip + " \\\n"
		command += "--ttl 10 -o table"
		shellExecOut(command)

		if oldIP != "" && oldIP != ip {
			cfmt.Info("Removing old IP:", ip)

			command = "az network dns record-set a remove-record \\\n"
			command += "-g tld \\\n"
			command += "-z " + ssl + " \\\n"
			command += "-n " + cluster + " \\\n"
			command += "-a " + oldIP + " -o table"
			shellExecOut(command)
		}
	}
}

// copy the files to the VM
func scpFilesToVM(ip string, ssl string) {
	cfmt.Info("Waiting for sshd service to start")

	// wait for sshd to start
	time.Sleep(30 * time.Second)

	// make sure we have permission to the directory
	sshCmd := "ssh -p 2222 -o \"StrictHostKeyChecking=no\" -o ConnectTimeout=600 akdc@" + ip + " \"sudo chown -R akdc:akdc /home/akdc\""
	shellExec(sshCmd)

	cfmt.Info("Copying customization files")
	scpToVM(ip, getParentDir()+"/vm-setup-files/akdc", "/home", 30, true)
	scpToVM(ip, "~/.ssh/akdc.pat", "~/.ssh/akdc.pat", 30, false)

	if ssl != "" {
		scpToVM(ip, "~/.ssh/certs.pem", "~/.ssh/certs.pem", 30, false)
		scpToVM(ip, "~/.ssh/certs.key", "~/.ssh/certs.key", 30, false)
	}

	echoStatusToVM(ip, "customization files copied")
}

// add a status message to the VM ~/status file
func echoStatusToVM(ip string, msg string) {
	shellExecOut("ssh -p 2222 -o \"StrictHostKeyChecking=no\" akdc@" + ip + " \"echo \"$(date +'%Y-%m-%d %H:%M:%S')  " + msg + "\" >> status\"")
}

// copy file(s) to VM worker
func scpToVM(ip string, source string, destination string, timeout int, recursive bool) {
	cmd := "scp -P 2222 -o \"StrictHostKeyChecking=no\" -o ConnectTimeout=" + strconv.Itoa(timeout) + " "

	if recursive {
		cmd += "-r "
	}

	cmd += source
	cmd += " akdc@" + ip + ":" + destination

	shellExec(cmd)
}
