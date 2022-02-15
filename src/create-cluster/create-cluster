#!/bin/bash

DirName=$(dirname "${BASH_SOURCE[0]}")
Templ="$DirName"/akdc.templ

# shellcheck source=/dev/null
source "${DirName}/lib/sh-arg.sh"

# shellcheck source=/dev/null
source "${DirName}/lib/logger.sh"

me=$(basename "$0")

# Input Configurations
DnsRecordTTL="1"
DnsZoneResourceGroup="tld"
DefaultLocation="centralus"
_argLength=$#

declare Location="$DefaultLocation"

# script arguments
declare DnsZone
declare UseSSL=false
declare ShowHelp=false
declare CertCRTPath=~/.ssh/certs.pem
declare CertKeyPath=~/.ssh/certs.key
declare Quiet=false
declare Repo=retaildevcrews/edge-gitops
declare Cluster
declare Group

# shellcheck disable=2034
declare DEBUG_FLAG=false

# register arguments
shArgs.arg "Cluster" -c --cluster PARAMETER true
shArgs.arg "Group" -g --group PARAMETER true
shArgs.arg "DnsZone" -z --zone PARAMETER true
shArgs.arg "CertCRTPath" -p --pem PARAMETER true
shArgs.arg "CertKeyPath" -k --key PARAMETER true
shArgs.arg "Location" -l --location PARAMETER true
shArgs.arg "Repo" -r --repo PARAMETER true
shArgs.arg "UseSSL" -s --ssl FLAG true
shArgs.arg "Quiet" -q --quiet FLAG true
shArgs.arg "ShowHelp" -h --help FLAG true
shArgs.arg "DEBUG_FLAG" -d --debug FLAG true

# parse inputs
shArgs.parse "$@"

FQDN="$Cluster.$DnsZone"

main() {
	if [ "$Quiet" != "true" ]; then
		banner
	fi

	validate_input

	if [ "$Quiet" != "true" ]; then
		print_context
	fi

	_information "Creating VM ${Cluster}"

	# create the RG tags
	local rgTags="server=$Cluster"

	if [ -n "$DnsZone" ]; then
		rgTags="$rgTags zone=$DnsZone"
	fi

	# create the install script from the template
	# replace the cluster, repo, pat, and fqdn
	rm -f "cluster-${Cluster}.sh"

	sed "s/{{pat}}/$AKDC_PAT/g" "$Templ" | \
		sed "s/{{cluster}}/$Cluster/g" | \
		sed "s/{{fqdn}}/$FQDN/g" | \
		sed "s~{{repo}}~$Repo~g" \
		> "cluster-${Cluster}.sh"

	# shellcheck disable=2086
	az group create -l "$Location" -n "$Group" --tags $rgTags -o table

	IP=$(az vm create \
		-g "$Group" \
		-l "$Location" \
		-n "$Cluster" \
		--admin-username akdc \
		--size standard_D2as_v5 \
		--image Canonical:0001-com-ubuntu-server-focal:20_04-lts-gen2:latest \
		--os-disk-size-gb 128 \
		--storage-sku Premium_LRS \
		--generate-ssh-keys \
		--public-ip-sku Standard \
		--query publicIpAddress \
		-o tsv \
		--custom-data "cluster-$Cluster.sh")

	echo -e "$Cluster\t$IP" >> ips

	_information "VM Created $Cluster  $IP"

	_information "Deleting default SSH rule"
	az network nsg rule delete -g "$Group" --nsg-name "${Cluster}"NSG -o table --name default-allow-ssh

	# For more security, replace --source-address-prefixes * with your IP or CIDR

	_information "Creating SSH rule on port 2222"
	az network nsg rule create \
		-g "$Group" \
		--nsg-name "${Cluster}"NSG \
		-n SSH-http \
		--description "SSH http https" \
		--destination-port-ranges 2222 80 443 \
		--protocol tcp \
		--access allow \
		--priority 1202 -o table

	if [ -n "$DnsZone" ]; then
		_information "Configuring DNS"
		_information "DnsZone: $DnsZone Name: $Cluster"
		upsert_dns_record
		
		if [ "$UseSSL" == "true" ]; then
			_information "Copying SSL Cert Files to VM"
			copy_cert_files_to_vm
		fi
	fi

	# validate ssh connectivity
	_information "connecting to VM with ssh"
	ssh -p 2222 -o "StrictHostKeyChecking=no" "akdc@$IP" 'cloud-init status'
	_success "VM Creation Complete! $Cluster  $IP"

	rm -f "cluster-$Cluster.sh"
}

validate_input() {
	if [ "$ShowHelp" == "true" ]; then
		usage
	fi

	if [ -z "$AKDC_PAT" ]; then
		_error "export AKDC_PAT with a valid Personal Access Token"
		usage
	fi

	if [ -z "$Cluster" ]; then
		_error "You must specify the cluster name with --cluster or -c"
		usage
	fi

	# default Azure RG to Cluster
	if [ -z "$Group" ]; then
		Group=$Cluster
	fi

	if [ "$UseSSL" == "true" ]; then
		if [ -z "$DnsZone" ]; then
			_error "DNS Zone (--zone) is a required argument when using --ssl"
			usage
		fi
	fi
}

wait_for_cloud_init_completion() {
	# this isn't currently used
	local ip=$1
	_information "Waiting for cloud init to complete."

	status=$(ssh -p 2222 -o "StrictHostKeyChecking=no" -o ConnectTimeout=60 akdc@"$ip" 'cloud-init status')
	_debug "cloud init status:$status."

	while [ "$status" != "status: done" ]; do
		sleep 10
		status=$(ssh -p 2222 -o "StrictHostKeyChecking=no" akdc@"$ip" 'cloud-init status')
		_debug "cloud init status:$status."
	done
}

tool_installed() {
	local tool="$1"
	local statusMessage

	if [ -x "$(command -v "$tool")" ]; then
			statusMessage="(installed - you're good to go!)"
	else
			statusMessage="\e[31m(not installed)\e[0m"
	fi

	echo "$statusMessage"
}

usage() {
		local jqStatusMessage && jqStatusMessage=$(tool_installed "jq")
		local sedStatusMessage && sedStatusMessage=$(tool_installed "sed")

		_helpText="  Usage: $me --cluster myClusterName [options][flags]

	Options and Flags
		-c | --cluster   <K8sCluster Name>     Kubernetes Cluster Name (required)
		-l | --location  <Azure Location>      Azure Location (default: centralus)
		-r | --repo      <GitOps Repo Name>    The GitOps repo name (default: retaildevcrews/edge-gitops)
		-z | --zone      <dns domain>          The dns domain to be configured on this VM
		-s | --ssl                             Configure an SSL cert for this domain. PEM and Key can be provided
		-p | --pem       <path to .pem file>   File Path to SSL Cert .pem file (default: ~/.ssh/certs.pem)
		-k | --key       <path to .key file>   File Path to SSL Cert .key file (default: ~/.ssh/certs.key)

	dependencies:
		-jq $jqStatusMessage
		-sed $sedStatusMessage"

		_information "$_helpText" 1>&2

		exit 1
}

banner() {
	cat <<- EOF
	 █████  ██   ██ ██████   ██████
	██   ██ ██  ██  ██   ██ ██
	███████ █████   ██   ██ ██
	██   ██ ██  ██  ██   ██ ██
	██   ██ ██   ██ ██████   ██████
	EOF

	echo "  version: $(cat "${DirName}"/version.txt)"
	echo "  k3d in VMs"
	echo ""
}

copy_cert_files_to_vm() {
	# set a long timeout so sshd has time to start
	scp -P 2222 -o "StrictHostKeyChecking=no" -o ConnectTimeout=600 "$CertCRTPath" akdc@"$IP":~/.ssh/certs.pem
	scp -P 2222 -o "StrictHostKeyChecking=no" "$CertKeyPath" akdc@"$IP":~/.ssh/certs.key
}

upsert_dns_record() {
	 OLD_IP=$(az network dns record-set a list \
	--query "[?name=='$Cluster'].{IP:aRecords}" \
	--resource-group "$DnsZoneResourceGroup" \
	--zone-name "$DnsZone"  -o json | jq -r '.[].IP[].ipv4Address')

	az network dns record-set a add-record \
		-g "$DnsZoneResourceGroup" \
		-z "$DnsZone" \
		-n "$Cluster" \
		-a "$IP" \
		--ttl $DnsRecordTTL -o table

	if [[ -n "$OLD_IP" ]]; then
		_information "Removing old IP: $OLD_IP from record $Cluster"

		az network dns record-set a remove-record \
			-g "$DnsZoneResourceGroup" \
			-z "$DnsZone" \
			-n "$Cluster" \
			-a "$OLD_IP" -o table
	fi
}

print_context() {
	_information "  @az Context"
	subscriptionId=$(az account show -o json | jq -r '.id')
	tenantId=$(az account show -o json | jq -r '.tenantId')
	subsscriptionName=$(az account show -o json | jq -r '.name')
	userName=$(az account show -o json | jq -r '.user.name')

	echo "  Subscription Id: $subscriptionId"
	echo "     Subscription: $subsscriptionName"
	echo "        Tenant Id: $tenantId"
	echo "             User: $userName"

	_information "  @Cluster Context"
	echo "            Cluster: $Cluster"
	echo "               FQDN: $FQDN"
	echo "     Azure Location: $Location"
	echo "     Resource Group: $Group"
	echo "        GitOps Repo: $Repo"
	echo "            DnsZone: $DnsZone"
	echo " DNS Resource Group: $DnsZoneResourceGroup"
	echo "            DNS TTL: $DnsRecordTTL"
	echo "            Use SSL: $UseSSL"
	echo "           SSL .pem: $CertCRTPath"
	echo "           SSL .key: $CertKeyPath"
}

main
