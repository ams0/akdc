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

var (
	// variables for options
	cluster           string
	group             string
	location          string
	repo              string
	branch            string
	pem               string
	key               string
	quiet             bool
	ssl               string
	dnsRG             string
	gitops            bool
	dapr              bool
	arcEnabled        bool
	digitalOcean      bool
	dryRun            bool
	debug             bool
	cores             int
	managedIdentityID = os.Getenv("AKDC_MI")

	// kic fleet create command
	CreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new cluster",
		Args:  validateCreateCmd,
		RunE:  runCreateCmd,
	}
)

// add kic fleet create specific flags
func init() {
	CreateCmd.Flags().StringVarP(&cluster, "cluster", "c", "", "Kubernetes cluster name (required)")
	CreateCmd.MarkFlagRequired("cluster")
	CreateCmd.Flags().StringVarP(&group, "group", "g", "", "Azure resource group name")
	CreateCmd.Flags().StringVarP(&location, "location", "l", "centralus", "Azure location")
	CreateCmd.Flags().StringVarP(&repo, "repo", "r", "retaildevcrews/edge-gitops", "GitOps repo name")
	CreateCmd.Flags().StringVarP(&branch, "branch", "b", "main", "GitOps branch name")
	CreateCmd.Flags().StringVarP(&ssl, "ssl", "s", "", "SSL domain name")
	CreateCmd.Flags().StringVarP(&pem, "pem", "p", "~/.ssh/certs.pem", "Path to SSL .pem file")
	CreateCmd.Flags().StringVarP(&key, "key", "k", "~/.ssh/certs.key", "Path to SSL .key file")
	CreateCmd.Flags().StringVarP(&dnsRG, "dns-resource-group", "", "tld", "DNS Resource Group")
	CreateCmd.Flags().BoolVarP(&dapr, "dapr", "", false, "Install Dapr and Radius")
	CreateCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Quiet mode")
	CreateCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Create VM in debug mode")
	CreateCmd.Flags().BoolVarP(&arcEnabled, "arc", "a", false, "Connect kubernetes cluster to Azure via Azure ARC")
	CreateCmd.Flags().BoolVarP(&digitalOcean, "do", "", false, "Generate setup script for Digital Ocean droplet")
	CreateCmd.Flags().BoolVarP(&gitops, "gitops", "", false, "Generate GitOps targets in --repo")
	CreateCmd.Flags().BoolVarP(&dryRun, "dry-run", "", false, "Show values that would be used")
	CreateCmd.Flags().IntVarP(&cores, "cores", "", 4, "VM core count")

	CreateCmd.RegisterFlagCompletionFunc("cluster", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getClusterComplete(), cobra.ShellCompDirectiveDefault
	})

	CreateCmd.RegisterFlagCompletionFunc("cores", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"02", "04", "08", "16", "32"}, cobra.ShellCompDirectiveDefault
	})
}

// get a list of valid clusters for shell completion
func getClusterComplete() []string {
	return []string{
		"central-tx-dallas-101",
		"central-tx-dallas-102",
		"central-tx-dallas-103",
		"central-tx-dallas-104",
		"central-tx-dallas-105",
		"east-ga-atl-101",
		"east-ga-atl-102",
		"east-ga-atl-103",
		"east-ga-atl-104",
		"east-ga-atl-105",
		"west-ca-la-101",
		"west-ca-la-102",
		"west-ca-la-103",
		"west-ca-la-104",
		"west-ca-la-105",
	}
}

// validation function for CreateCmd
func validateCreateCmd(cmd *cobra.Command, args []string) error {
	// validate ssl and domain
	hasError := validateSSL()

	// validate cluster name against reserved prefixes (if set)
	hasError = hasError && validateClusterPrefix()

	// managed identity is required for Azure VMs
	hasError = hasError && validateManagedIdentity()

	// validate cores
	hasError = hasError && validateCores()

	if hasError {
		return fmt.Errorf("create command aborted")
	}

	// default resource group is cluster name
	if group == "" {
		group = cluster
	}

	return nil
}

// validate --ssl
func validateSSL() bool {
	hasError := false

	if ssl != "" {
		if pem == "" {
			cfmt.ErrorE("you must specify --pem to use --ssl")
			hasError = true
		}

		if key == "" {
			cfmt.ErrorE("you must specify --key to use --ssl")
			hasError = true
		}

		if len(ssl) < 3 {
			cfmt.ErrorE("--ssl parameter is too short")
			hasError = true
		}

		if !strings.Contains(ssl, ".") {
			cfmt.ErrorE("malformed --ssl parameter")
			hasError = true
		}
	}
	return hasError
}

// validate --cores
func validateCores() bool {
	validCores := map[int]int{2: 2, 4: 4, 8: 8, 16: 16, 32: 32}
	_, ok := validCores[cores]

	if !ok {
		cfmt.ErrorE("invalid --cores")
		fmt.Println("  valid: 2, 4, 8, 16, 32")
	}

	return ok
}

// validate managed identity
func validateManagedIdentity() bool {
	if !digitalOcean && managedIdentityID == "" {
		cfmt.ErrorE("managed identity is required")
		fmt.Println("  export AKDC_MI=yourManagedIdentity")
		return false
	}

	return true
}

// validate cluster prefix
func validateClusterPrefix() bool {
	cl := strings.ToLower(cluster)

	blocked := boa.ReadConfigValue("reservedClusterPrefixes:")

	lines := strings.Split(blocked, " ")

	for _, line := range lines {
		line = strings.ToLower(strings.TrimSpace(line))

		if line != "" && strings.HasPrefix(cl, line) {
			cfmt.ErrorE("cluster name is invalid - reserved prefix")
			return false
		}
	}

	return true
}

// run the command
func runCreateCmd(cmd *cobra.Command, args []string) error {
	// create the setup script from the template
	createVMSetupScript()

	if dryRun {
		return doDryRun()
	}

	if digitalOcean {
		// add the GitOps target
		addTarget(cluster, ssl)

		// no more automation for Digital Ocean droplets
		cfmt.Info("Digital Ocean template created")
		return nil
	}

	// create the azure resource group
	if err := createGroup(); err != nil {
		cfmt.ErrorE(err)
		return err
	}

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
}

// handle --dry-run
func doDryRun() error {
	fmt.Println("Cluster:             ", cluster)
	fmt.Println("Cores:               ", cores)
	fmt.Println("Group:               ", group)

	fmt.Println("Managed Identity:    ", strings.Contains(managedIdentityID, "/Microsoft.ManagedIdentity/"))
	fmt.Println("Location:            ", location)
	fmt.Println("Repo:                ", repo)
	fmt.Println("Branch:              ", branch)

	if len(ssl) > 0 {
		fmt.Println("SSL Domain:          ", ssl)
		fmt.Println("SSL pem:             ", pem)
		fmt.Println("SSL key:             ", key)
	} else {
		fmt.Println("SSL Domain:           none")
	}

	fmt.Println("Enable Arc:          ", arcEnabled)
	fmt.Println("Enable Dapr:         ", dapr)
	fmt.Println("Enable Digital Ocean:", digitalOcean)
	fmt.Println("Enable GitOps:       ", gitops)

	if !digitalOcean {
		fmt.Println("Cluster Exists:      ", vmExists())
		fmt.Println("Group Exists:        ", groupExists())
	}

	return nil
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

// get the path to template file
func getTemplatePath() string {
	return boa.GetRepoBase() + "/assets/akdc.templ"
}

// create Azure Resource Group
func createGroup() error {
	// fail if the Azure VM exists
	if vmExists() {
		cfmt.ErrorE("Azure VM exists")
		return fmt.Errorf("Please use a different VM or delete the VM")
	}

	// don't create the group if it exists
	if groupExists() {
		return nil
	}

	cfmt.Info("Creating Azure Resource Group")

	rgTags := "akdc=true server=" + cluster

	if ssl != "" {
		rgTags += " zone=" + ssl
	}

	command := "az group create -l " + location + " -n " + group + " -o table --tags " + rgTags

	err := boa.ShellExecE(command)
	if err != nil {
		cfmt.ErrorE(err)
	}

	return err
}

// create vm setup script
func createVMSetupScript() {
	os.Remove("cluster-" + cluster + ".sh")

	// create the custom VM script
	content, err := os.ReadFile(getTemplatePath())

	if err != nil {
		cfmt.ErrorE(err)
		return
	}

	// create the vm setup script from the template
	command := string(content)
	command = strings.Replace(command, "{{cluster}}", cluster, -1)
	command = strings.Replace(command, "{{dapr}}", strconv.FormatBool(dapr), -1)
	command = strings.Replace(command, "{{debug}}", strconv.FormatBool(debug), -1)
	command = strings.Replace(command, "{{fqdn}}", cluster+"."+ssl, -1)
	command = strings.Replace(command, "{{repo}}", repo, -1)
	command = strings.Replace(command, "{{branch}}", branch, -1)
	command = strings.Replace(command, "{{group}}", group, -1)
	command = strings.Replace(command, "{{arcEnabled}}", strconv.FormatBool(arcEnabled), -1)
	command = strings.Replace(command, "{{do}}", strconv.FormatBool(digitalOcean), -1)
	command = strings.Replace(command, "{{zone}}", ssl, -1)
	command = strings.Replace(command, "{{dnsRG}}", dnsRG, -1)
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
	command += " --assign-identity " + managedIdentityID + "\\\n"
	command += " --size standard_D" + strconv.Itoa(cores) + "as_v5 \\\n"
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
		cfmt.ErrorE("Failed to create cluster")
		return ""
	}
	cfmt.Info("VM created")
	fmt.Println(cluster, ip)

	f, err := os.OpenFile("ips", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		cfmt.ErrorE(err)
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

// get GitOps template
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
		cfmt.ErrorE("unable to read gitops-config.templ")
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

		gitopsDir := filepath.Join("/", "workspaces")

		if _, err := os.Stat(gitopsDir); err == nil {
			gitopsDir = filepath.Join(gitopsDir, repoName)
		} else {
			gitopsDir = filepath.Join(os.Getenv("HOME"), repoName)
		}

		bootstrapDir := filepath.Join(gitopsDir, "deploy", "bootstrap")
		appsDir := filepath.Join(gitopsDir, "deploy", "apps")

		dirExists := false

		gitCmd := "git -C " + gitopsDir + " "

		// clone or pull the repo
		if _, err := os.Stat(gitopsDir); err == nil {
			dirExists = true
		} else {
			boa.ShellExecE("git clone " + repoFull + " " + gitopsDir)
		}

		// checkout the branch
		boa.ShellExecE(gitCmd + "pull")
		boa.ShellExecE(gitCmd + "checkout " + branch)
		boa.ShellExecE(gitCmd + "pull")

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
