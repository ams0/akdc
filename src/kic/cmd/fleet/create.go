// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"fmt"
	"kic/boa"
	"kic/boa/cfmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// mainly for create command
var cluster string
var group string
var location string
var repo string
var pem string
var key string
var quiet bool
var ssl string
var gitops bool

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

		blocked := boa.ReadConfigValue("reservedClusterPrefixes:")

		lines := strings.Split(blocked, " ")

		for _, line := range lines {
			line = strings.ToLower(strings.TrimSpace(line))

			if line != "" && strings.HasPrefix(cluster, line) {
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

	RunE: func(cmd *cobra.Command, args []string) error {
		// create the setup script from the template
		createVMSetupScript()

		if digitalOcean {
			// add the GitOps target
			addTarget(cluster, ssl)

			// no more automation for Digital Ocean droplets
			cfmt.Info("Digital Ocean template created")
			return nil
		}

		managedIdentityID := os.Getenv("AKDC_MI")
		if managedIdentityID == "" {
			cfmt.Error("managed identity is required")
			fmt.Println("  export AKDC_MI=yourManagedIdentity")
			return fmt.Errorf("VM Creation Failed")
		}
		// create the azure resource group
		createGroup()

		// create the vm and get the IP
		ip := createVM(managedIdentityID)

		// remove the cluster template
		os.Remove("cluster-" + cluster + ".sh")

		// success
		if ip != "" {
			// add the GitOps target
			addTarget(cluster, ssl)

			cfmt.Info("VM Configured")
			return nil
		}

		return fmt.Errorf("VM Creation Failed")
	},
}

// check to see if the Azure Resource Group exists
func groupExists() bool {
	ex, _ := boa.ShellExecOut("az group exists -g " + group)
	return strings.HasPrefix(ex, "true")
}

// check to see if the VM exists in the RG
func vmExists() bool {
	command := fmt.Sprintf("az vm show -g %s --name %s --query 'name' -o tsv", group, cluster)
	res, _ := boa.ShellExecOut(command)
	return strings.EqualFold(cluster, strings.TrimSpace(res))
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
	CreateCmd.Flags().BoolVarP(&gitops, "gitops", "", false, "Generate GitOps targets in --repo")
}

// get the path to template file
func getTemplatePath() string {
	return boa.GetRepoBase() + "/assets/akdc.templ"
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

	err := boa.ShellExecE(command)
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

	ip, _ := boa.ShellExecOut(command)
	ip = strings.TrimSpace(ip)

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
	boa.ShellExecOut(command)

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
	boa.ShellExecOut(command)

	return ip
}

// create DNS entry and copy SSL certs
func getConfigJson(cluster string, zone string) []byte {
	region := cluster
	district := cluster

	cols := strings.Split(cluster, "-")

	if len(cols) > 0 {
		region = cols[0]

		if len(cols) > 2 {
			district = strings.Join(cols[:3], "-")
		}
	}

	json := boa.ReadTextFileFromBoaDir("gitops-config.templ")

	if json == "" {
		return nil
	}

	// replace template values
	json = strings.Replace(json, "{{cluster}}", cluster, -1)
	json = strings.Replace(json, "{{region}}", region, -1)
	json = strings.Replace(json, "{{district}}", district, -1)
	json = strings.Replace(json, "{{zone}}", zone, -1)

	return []byte(json)
}

// add a target to GitOps
func addTarget(cluster string, zone string) {
	// only run if --gitops specified
	if !gitops {
		return
	}

	// read the config.json file
	json := getConfigJson(cluster, ssl)

	if json == nil {
		cfmt.Error("unable to read gitops-config.templ")
	} else {
		// GitOps directories
		repoName := repo
		repoFull := repoName

		if strings.HasPrefix(repo, "https://") {
			cols := strings.Split(repo[8:], "/")
			repoName = strings.Join(cols[len(cols)-2:], "/")
		} else {
			repoFull = "https://github.com/" + repo
		}

		cols := strings.Split(repoName, "/")
		repoName = cols[len(cols)-1]

		gitopsDir := filepath.Join(os.Getenv("HOME"), repoName)
		bootstrapDir := filepath.Join(gitopsDir, "deploy", "bootstrap")
		appsDir := filepath.Join(gitopsDir, "deploy", "apps")

		dirExists := false

		gitCmd := "git -C " + gitopsDir + " "

		// clone or pull the repo
		if _, err := os.Stat(gitopsDir); err == nil {
			boa.ShellExecE(gitCmd + "pull")
			dirExists = true
		} else {
			boa.ShellExecE("git clone " + repoFull + " " + gitopsDir)
		}

		// add the targets
		if len(json) > 0 && !strings.Contains(string(json), "{{") {
			// make sure the dirs exist
			if _, err := os.Stat(bootstrapDir); err == nil {
				if _, err = os.Stat(appsDir); err == nil {
					// add cluster to the dirs
					bootstrapDir = filepath.Join(bootstrapDir, cluster)
					appsDir = filepath.Join(appsDir, cluster)

					// create the directories
					boa.ShellExecE("mkdir -p " + bootstrapDir + " && mkdir -p " + appsDir)

					// write config.json to each dir
					bootstrapDir = filepath.Join(bootstrapDir, "config.json")
					appsDir = filepath.Join(appsDir, "config.json")

					// don't overwrite existing config.json
					if _, err := os.Stat(bootstrapDir); err != nil {
						os.WriteFile(bootstrapDir, json, 0644)
					}

					if _, err := os.Stat(appsDir); err != nil {
						os.WriteFile(appsDir, json, 0644)
					}

					// if there were repo changes
					if res, _ := boa.ShellExecOut(gitCmd + "status -s"); res != "" {
						agoDir := filepath.Join(gitopsDir, "autogitops")

						if _, err := os.Stat(agoDir); err == nil {
							if file, err := os.OpenFile(filepath.Join(agoDir, "create.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err == nil {
								dt := time.Now().UTC()
								fmt.Fprintln(file, dt.Format("01-02-2006 15:04:05"), "added cluster:", cluster)
								file.Close()
							}
						}
						// update the repo
						boa.ShellExecE(gitCmd + "add .")
						boa.ShellExecE(gitCmd + "commit -am 'kic fleet create'")
						boa.ShellExecE(gitCmd + "pull")
						boa.ShellExecE(gitCmd + "push")
					}

					// delete the repo if we created it
					if !dirExists {
						boa.ShellExecE("rm -rf " + gitopsDir)
					}
				}
			}
		}
	}
}
