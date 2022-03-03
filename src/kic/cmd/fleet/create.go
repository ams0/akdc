// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"fmt"
	"kic/cfmt"
	"kic/utils"
	"log"
	"os"

	// "path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new cluster",

	Args: func(cmd *cobra.Command, args []string) error {
		hasError := false

		// validate ssl and domain
		if ssl != "" {
			if pem == "" {
				cfmt.Error("you must specify --pem to use --ssl")
				hasError = true
			}

			if key == "" {
				cfmt.Error("you must specify --key to use --ssl")
				hasError = true
			}

			if len(ssl) < 3 {
				cfmt.Error("--ssl parameter is too short")
				hasError = true
			}

			if !strings.Contains(ssl, ".") {
				cfmt.Error("malformed --ssl parameter")
				hasError = true
			}
		}

		cluster = strings.ToLower(cluster)

		blocked := utils.ReadConfigValue("reservedClusterPrefixes:")

		lines := strings.Split(blocked, " ")

		for _, line := range lines {
			line = strings.ToLower(strings.TrimSpace(line))

			if strings.HasPrefix(cluster, line) {
				cfmt.Error("cluster name is invalid - reserved prefix")
				hasError = true
			}
		}

		// default resource group is cluster name
		if group == "" {
			group = cluster
		}

		if hasError {
			return fmt.Errorf("create command aborted")
		}

		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		// create the setup script from the template
		createVMSetupScript()

		if digitalOcean {
			cfmt.Info("Digital Ocean template created")
			// no more automation for Digit Ocean droplets
			return
		}

		managedIdentityID := os.Getenv("AKDC_MI")
		if managedIdentityID == "" {
			cfmt.Error("managed identity is required")
			fmt.Println("  export AKDC_MI=yourManagedIdentity")
			os.Exit(1)
		}
		// create the azure resource group
		createGroup()

		// create the vm and get the IP
		ip := createVM(managedIdentityID)

		// fail if createVM fails
		if ip != "" {
			// create DNS entry
			createDNS(ip)
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
	ex := utils.ShellExecOut("az group exists -g " + group)
	return strings.HasPrefix(ex, "true")
}

// check to see if the VM exists in the RG
func vmExists() bool {
	command := fmt.Sprintf("az vm show -g %s --name %s --query 'name' -o tsv", group, cluster)
	res := strings.TrimSpace(utils.ShellExecOut(command))
	return strings.EqualFold(cluster, res)
}

var dapr bool
var arcEnabled bool
var digitalOcean bool

// add akdc create specific flags
func init() {
	CreateCmd.Flags().StringVarP(&cluster, "cluster", "c", "", "Kubernetes cluster name (required)")
	CreateCmd.MarkFlagRequired("cluster")
	CreateCmd.Flags().StringVarP(&group, "group", "g", "", "Azure resource group name")
	CreateCmd.Flags().StringVarP(&location, "location", "l", "centralus", "Azure location")
	CreateCmd.Flags().StringVarP(&repo, "repo", "r", "retaildevcrews/edge-gitops", "GitOps repo name")
	CreateCmd.Flags().StringVarP(&ssl, "ssl", "s", "", "SSL domain name")
	CreateCmd.Flags().StringVarP(&pem, "pem", "p", "~/.ssh/certs.pem", "Path to SSL .pem file")
	CreateCmd.Flags().StringVarP(&key, "key", "k", "~/.ssh/certs.key", " Path to SSL .key file")
	CreateCmd.Flags().BoolVarP(&dapr, "dapr", "", false, "Install Dapr and Radius")
	CreateCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Quiet mode")
	CreateCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Create VM in debug mode")
	CreateCmd.Flags().BoolVarP(&arcEnabled, "arc", "a", false, "Connect kubernetes cluster to Azure via Azure ARC")
	CreateCmd.Flags().BoolVarP(&digitalOcean, "do", "", false, "Generate setup script for Digital Ocean droplet")
}

// get the path to template file
func getTemplatePath() string {
	return utils.GetRepoBase() + "/assets/akdc.templ"
}

// create Azure Resource Group
func createGroup() {
	// fail if the Azure VM exists
	if vmExists() {
		cfmt.Error("Azure VM exists")
		cfmt.ExitErrorMessage("Please use a different VM or delete the VM")
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

	err := utils.ShellExecE(command)
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
	command = strings.Replace(command, "{{do}}", strconv.FormatBool(digitalOcean), -1)
	command = strings.Replace(command, "{{zone}}", ssl, -1)
	// todo - define dnsRG
	// command = strings.Replace(command, "{{dnsRG}}", dnsRG, -1)
	os.WriteFile("cluster-"+cluster+".sh", []byte(command), 0644)
}

// create Azure VM
func createVM(managedIdentityID string) string {
	cfmt.Info("Creating Azure VM")

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

	ip := strings.TrimSpace(utils.ShellExecOut(command))

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
	utils.ShellExecOut(command)

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
	utils.ShellExecOut(command)

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

		oldIP := strings.TrimSpace(utils.ShellExecOut(command))

		command = "az network dns record-set a add-record \\\n"
		command += "-g tld \\\n"
		command += "-z " + ssl + " \\\n"
		command += "-n " + cluster + " \\\n"
		command += "-a " + ip + " \\\n"
		command += "--ttl 10 -o table"
		utils.ShellExecOut(command)

		if oldIP != "" && oldIP != ip {
			cfmt.Info("Removing old IP:", ip)

			command = "az network dns record-set a remove-record \\\n"
			command += "-g tld \\\n"
			command += "-z " + ssl + " \\\n"
			command += "-n " + cluster + " \\\n"
			command += "-a " + oldIP + " -o table"
			utils.ShellExecOut(command)
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
	utils.ShellExec(sshCmd)

	cfmt.Info("Copying customization files")
	scpToVM(ip, "~/.ssh/akdc.pat", "~/.ssh/akdc.pat", 30, false)

	if ssl != "" {
		scpToVM(ip, "~/.ssh/certs.pem", "~/.ssh/certs.pem", 30, false)
		scpToVM(ip, "~/.ssh/certs.key", "~/.ssh/certs.key", 30, false)
	}

	echoStatusToVM(ip, "customization files copied")
}

// add a status message to the VM ~/status file
func echoStatusToVM(ip string, msg string) {
	utils.ShellExecOut("ssh -p 2222 -o \"StrictHostKeyChecking=no\" akdc@" + ip + " \"echo \"$(date +'%Y-%m-%d %H:%M:%S')   " + msg + "\" >> status\"")
}

// copy file(s) to VM worker
func scpToVM(ip string, source string, destination string, timeout int, recursive bool) {
	cmd := "scp -P 2222 -o \"StrictHostKeyChecking=no\" -o ConnectTimeout=" + strconv.Itoa(timeout) + " "

	if recursive {
		cmd += "-r "
	}

	cmd += source
	cmd += " akdc@" + ip + ":" + destination

	utils.ShellExec(cmd)
}
