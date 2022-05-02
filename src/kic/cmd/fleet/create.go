// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"fmt"
	"kic/boa"
	"kic/boa/cfmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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
	gitopsOnly        bool
	dapr              bool
	arcEnabled        bool
	digitalOcean      bool
	dryRun            bool
	debug             bool
	cores             int
	verbose           bool
	sku               string
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
	// get defaults from env var
	// todo - convert to Viper
	tRepo := os.Getenv("AKDC_REPO")
	if tRepo == "" {
		tRepo = "retaildevcrews/edge-gitops"
	}

	tSsl := os.Getenv("AKDC_SSL")

	tGitOps := os.Getenv("AKDC_GITOPS") == "true"

	CreateCmd.Flags().StringVarP(&cluster, "cluster", "c", "", "Kubernetes cluster name (required)")
	CreateCmd.MarkFlagRequired("cluster")
	CreateCmd.Flags().StringVarP(&group, "group", "g", "", "Azure resource group name")
	CreateCmd.Flags().StringVarP(&location, "location", "l", "centralus", "Azure location")
	CreateCmd.Flags().StringVarP(&repo, "repo", "r", tRepo, "GitOps repo name")
	CreateCmd.Flags().StringVarP(&branch, "branch", "b", "", "GitOps branch name")
	CreateCmd.Flags().StringVarP(&ssl, "ssl", "s", tSsl, "SSL domain name")
	CreateCmd.Flags().StringVarP(&pem, "pem", "p", "~/.ssh/certs.pem", "Path to SSL .pem file")
	CreateCmd.Flags().StringVarP(&key, "key", "k", "~/.ssh/certs.key", "Path to SSL .key file")
	CreateCmd.Flags().StringVarP(&dnsRG, "dns-resource-group", "", "tld", "DNS Resource Group")
	CreateCmd.Flags().BoolVarP(&dapr, "dapr", "", false, "Install Dapr and Radius")
	CreateCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Quiet mode")
	CreateCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Create VM in debug mode")
	CreateCmd.Flags().BoolVarP(&arcEnabled, "arc", "a", false, "Connect kubernetes cluster to Azure via Azure ARC")
	CreateCmd.Flags().BoolVarP(&digitalOcean, "do", "", false, "Generate setup script for Digital Ocean droplet")
	CreateCmd.Flags().BoolVarP(&gitops, "gitops", "", tGitOps, "Generate GitOps targets in ./config")
	CreateCmd.Flags().BoolVarP(&gitopsOnly, "gitops-only", "", false, "Only generate GitOps targets in ./config")
	CreateCmd.Flags().BoolVarP(&dryRun, "dry-run", "", false, "Show values that would be used")
	CreateCmd.Flags().BoolVarP(&verbose, "verbose", "", false, "Show verbose output")
	CreateCmd.Flags().IntVarP(&cores, "cores", "", 4, "VM core count")
	CreateCmd.Flags().StringVarP(&sku, "sku", "", "", "Azure VM SKU")

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

// validate logged in to Azure
func validateAzureLogin() error {
	_, err := boa.ShellExecOut("az account show --query tenantId -o tsv", false)
	return err
}

// validate Azure region
func validateLocation(loc string) error {
	res, err := boa.ShellExecOut("az account list-locations --query '[].name' -o table |  grep -x -i "+location, false)

	if err != nil || res == "" {
		return fmt.Errorf("%s", "invalid Azure Region")
	}

	return nil
}

// validate VM SKU
func validateVmSku(loc string, vmSku string) error {
	res, err := boa.ShellExecOut("az vm list-sizes -o table -l "+location+" | grep -w -i "+vmSku, false)

	if err != nil || res == "" {
		return fmt.Errorf("%s", "invalid VM SKU")
	}

	return nil
}

// validate PWD is a git repo
func validateRepo() bool {
	// required for --gitops
	if gitops || gitopsOnly {
		res, err := boa.ShellExecOut("git branch --show-current", false)

		if err != nil {
			cfmt.ErrorE("Not a git repo")
			cfmt.Info("Please re-run from a git repo")
			return false
		}

		res = strings.TrimSpace(res)

		// set branch to current branch
		if branch == "" {
			branch = res
		}
	}

	return true
}

// run the command
func runCreateCmd(cmd *cobra.Command, args []string) error {
	if !digitalOcean {
		// validate PWD is a git repo
		if !validateRepo() {
			return nil
		}

		// validate azure login
		if validateAzureLogin() != nil {
			cfmt.ErrorE("please run az login first")
			return nil
		}
		// validate location
		location = strings.ToLower(location)

		if err := validateLocation(location); err != nil {
			cfmt.ErrorE("Invalid location")
			cfmt.Info("Valid Locations")
			boa.ShellExecE("az account list-locations --query '[].name' -o table | sort")
			fmt.Println()
			cfmt.ErrorE("Invalid location")
			return nil
		}

		// SKU was not specified - try defaults
		if sku == "" {
			// check the D4as_V5 SKU
			sku = "Standard_D" + strconv.Itoa(cores) + "as_v5"

			err := validateVmSku(location, sku)

			if err != nil {
				// check the D4s_v5 SKU
				sku = "Standard_D" + strconv.Itoa(cores) + "s_v5"
			}
		}

		// validate VM SKU
		if err := validateVmSku(location, sku); err != nil {
			cfmt.ErrorE("Invalid SKU")
			cfmt.Info("Valid SKUs for region: " + location)
			boa.ShellExecE("az vm list-sizes -l " + location + " -o table | grep _v5 | grep Standard_D" + " |awk '{print $3}' | sort")
			fmt.Println()
			cfmt.ErrorE("Invalid SKU")
			return nil
		}
	}

	if dryRun {
		return doDryRun()
	}

	// add the GitOps target
	if gitops || gitopsOnly {
		addTargetE(cluster)
		cfmt.Info("Created GitOps config: ", cluster)

		// exit
		if gitopsOnly {
			return nil
		}
	}

	// fail if the Azure VM exists
	if vmExists() {
		cfmt.ErrorE("Azure VM Exists: ", cluster)
		return nil
	}

	// create the setup script from the template
	createVMSetupScript()

	if digitalOcean {
		// no more automation for Digital Ocean droplets
		cfmt.Info("Digital Ocean template created")
		return nil
	}

	// create the azure resource group
	if err := createGroup(); err != nil {
		cfmt.ErrorE("createGroup Failed: ", cluster)
		return nil
	}

	// create the vm and get the IP
	ip := createVM(managedIdentityID)

	// remove the cluster template
	os.Remove("cluster-" + cluster + ".sh")

	// success
	if ip != "" {
		cfmt.Info("VM Configured: ", cluster)
	} else {
		cfmt.ErrorE("VM Creation Failed: ", cluster)
	}

	return nil
}

// handle --dry-run
func doDryRun() error {
	fmt.Println("Cluster:             ", cluster)
	fmt.Println("Cores:               ", cores)
	fmt.Println("Group:               ", group)

	if !digitalOcean {
		fmt.Println("VM SKU:              ", sku)
	}

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
	ex, _ := boa.ShellExecOut("az group exists -g "+group, false)
	return strings.HasPrefix(ex, "true")
}

// check to see if the VM exists in the RG
func vmExists() bool {
	command := fmt.Sprintf("az vm show -g %s --name %s --query 'name' -o tsv", group, cluster)
	res, _ := boa.ShellExecOut(command, false)
	return strings.EqualFold(cluster, strings.TrimSpace(res))
}

// get the path to template file
func getTemplatePath() string {
	return boa.GetBoaPath() + "/fleet-vm.templ"
}

// create Azure Resource Group
func createGroup() error {
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
	command = strings.ReplaceAll(command, "{{cluster}}", cluster)
	command = strings.ReplaceAll(command, "{{dapr}}", strconv.FormatBool(dapr))
	command = strings.ReplaceAll(command, "{{debug}}", strconv.FormatBool(debug))
	command = strings.ReplaceAll(command, "{{fqdn}}", cluster+"."+ssl)
	command = strings.ReplaceAll(command, "{{repo}}", repo)
	command = strings.ReplaceAll(command, "{{branch}}", branch)
	command = strings.ReplaceAll(command, "{{group}}", group)
	command = strings.ReplaceAll(command, "{{arcEnabled}}", strconv.FormatBool(arcEnabled))
	command = strings.ReplaceAll(command, "{{do}}", strconv.FormatBool(digitalOcean))
	command = strings.ReplaceAll(command, "{{zone}}", ssl)
	command = strings.ReplaceAll(command, "{{dnsRG}}", dnsRG)

	// todo - testing
	env := ""

	for _, val := range os.Environ() {
		if strings.HasPrefix(val, "AKDC_") {
			split := strings.SplitN(val, "=", 2)
			if len(split) == 2 {
				val = strings.ReplaceAll(split[1], "\n", "")
				val = strings.ReplaceAll(val, "\"", "\\\"")
				line := fmt.Sprintf("  echo 'export %s=\"", split[0])
				line += val
				line += "\"'\n"

				env += line
			}
		}
	}
	command = strings.ReplaceAll(command, "{{environment}}", env)

	os.WriteFile("cluster-"+cluster+".sh", []byte(command), 0644)
}

// create Azure VM
func createVM(managedIdentityID string) string {
	cfmt.Info("Creating Azure VM: ", cluster)

	command := "az vm create \\\n"
	command += " -g " + group + " \\\n"
	command += " -l " + location + " \\\n"
	command += " -n " + cluster + " \\\n"
	command += " --admin-username akdc \\\n"
	command += " --assign-identity " + managedIdentityID + "\\\n"
	command += " --size " + sku + " \\\n"
	command += " --image Canonical:0001-com-ubuntu-server-focal:20_04-lts-gen2:latest \\\n"
	command += " --os-disk-size-gb 128 \\\n"
	command += " --storage-sku Premium_LRS \\\n"
	command += " --generate-ssh-keys \\\n"
	command += " --public-ip-sku Standard \\\n"
	command += " --custom-data cluster-" + cluster + ".sh \\\n"
	command += " --query publicIpAddress \\\n"
	command += " -o tsv"

	ip, err := boa.ShellExecOut(command, verbose)
	ip = strings.TrimSpace(ip)

	if err != nil || ip == "" {
		cfmt.FAppendToFile("failed.log", cluster+"\n")
		return ""
	}

	cfmt.Info("VM Created: ", cluster)
	fmt.Println(cluster, ip)
	cfmt.FAppendToFile("ips", fmt.Sprintf("%s\t%s\n", cluster, ip))

	cfmt.Info("Deleting NSG: ", cluster)
	command = "az network nsg rule delete -g " + group + " --nsg-name " + cluster + "NSG -o table --name default-allow-ssh"
	boa.ShellExecOut(command, false)

	cfmt.Info("Creating SSH Rule: ", cluster)

	command = "az network nsg rule create \\\n"
	command += "-g " + group + " \\\n"
	command += "--nsg-name " + cluster + "NSG \\\n"
	command += "-n SSH-http \\\n"
	command += "--description \"SSH http https\" \\\n"
	command += "--destination-port-ranges 2222 80 443 \\\n"
	command += "--protocol tcp \\\n"
	command += "--access allow \\\n"
	command += "--priority 1202 -o table"
	boa.ShellExecOut(command, false)

	return ip
}

// get GitOps template
func getConfigJson(cluster string) []byte {
	region := cluster
	zone := cluster
	district := cluster

	cols := strings.Split(cluster, "-")

	if len(cols) > 0 {
		region = cols[0]

		if len(cols) > 1 {
			zone = strings.Join(cols[0:2], "-")
		}

		if len(cols) > 2 {
			district = strings.Join(cols[0:3], "-")
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
	json = strings.Replace(json, "{{domain}}", ssl, -1)

	return []byte(json)
}

// add a target to GitOps
func addTargetE(cluster string) error {
	// only run if --gitops specified
	if !(gitops || gitopsOnly) {
		return nil
	}

	// read the config.json file
	json := getConfigJson(cluster)

	// make sure the json is valid
	if json == nil || len(json) < 3 || strings.Contains(string(json), "{{") {
		return cfmt.ErrorE("unable to read gitops-config.templ")
	}

	return addTargetWorkerE(json)
}

// add a target to GitOps
func addTargetWorkerE(json []byte) error {
	configDir := filepath.Join(".", "config")

	// add the targets
	if len(json) > 0 && !strings.Contains(string(json), "{{") {
		// make sure the dirs exist
		if _, err := os.Stat(configDir); err == nil {
			// add cluster to the dirs
			configDir = filepath.Join(configDir, cluster)

			// create the directory
			if err := boa.ShellExecE("mkdir -p " + configDir); err != nil {
				return err
			}

			// write config.json to each dir
			configDir = filepath.Join(configDir, "config.json")

			if _, err := os.Stat(configDir); err != nil {
				if err := os.WriteFile(configDir, json, 0644); err != nil {
					return err
				}
			}

			return updateGitOpsRepoE()
		}
	}

	return nil
}

// update the GitOps repo with changes
func updateGitOpsRepoE() error {
	// if there were repo changes
	if res, _ := boa.ShellExecOut("git status -s", false); res != "" {
		// pull to avoid conflicts
		if _, err := boa.ShellExecOut("git pull", false); err != nil {
			return err
		}

		// update the repo
		if _, err := boa.ShellExecOut("git add .", false); err != nil {
			return err
		}

		if _, err := boa.ShellExecOut("git commit -am 'flt create'", false); err != nil {
			return err
		}

		// push the changes
		if _, err := boa.ShellExecOut("git push", false); err != nil {
			return err
		}
	}

	return nil
}
