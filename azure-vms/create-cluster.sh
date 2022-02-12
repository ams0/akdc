#!/bin/bash

source "lib/sh-arg.sh"
source "lib/logger.sh"

me=`basename "$0"`

# Input Configurations
Region=${1}
State=${2}
City=${3}
Number=${4}
District=${Region}-${State}-${City}
Store=${District}-${Number}
DnsRecordTTL="1"
DnsZoneResourceGroup="tld"
DefaultLocation="centralus"
_argLength=$#


declare Location="$DefaultLocation"

# script arguments
declare DnsZone
declare UseSSL=false
declare ShowHelp=false
declare DEBUG_FLAG=false
declare CertCRTPath
declare CertKeyPath
# register arguments
shArgs.arg "DnsZone" -z --zone PARAMETER true
shArgs.arg "CertCRTPath" -c --crt PARAMETER true
shArgs.arg "CertKeyPath" -k --key PARAMETER true
shArgs.arg "Location" -l --location PARAMETER true
shArgs.arg "UseSSL" -s --ssl FLAG true 
shArgs.arg "ShowHelp" -h --usage FLAG true 
shArgs.arg "DEBUG_FLAG" -d --debug FLAG true

# parse inputs
shArgs.parse $@

FQDN="$Store.$DnsZone"

# location is a special case as it can be passed in as the fifth argument or as --location (via shArg)
if [[ "$5" != -* ]] ; then
  if [[ "$5" != "" ]]; then
    Location=${5}
  fi
fi

main() {
  banner
  validate_input
  print_context

  # create the RG
  local rgTags="server=$Store"
  if [ ! -z "$DnsZone" ]; then
    rgTags="$rgTags zone=$DnsZone"
  fi
  az group create -l $Location -n $Store --tags $rgTags

  # create the install script from the template
  # replace the host, pat, district and region
  rm -f cluster-$Store.sh
  sed "s/{{pat}}/$AKDC_PAT/g" ./akdc.templ | \
      sed "s/{{store}}/$Store/g" | \
      sed "s/{{district}}/$District/g" | \
      sed "s/{{fqdn}}/$FQDN/g" | \
      sed "s/{{region}}/$Region/g" \
      > cluster-$Store.sh

  _information "Creating VM"
  IP=$(az vm create \
    -g $Store \
    --admin-username akdc \
    -n $Store \
    --size standard_D2as_v5 \
    --image Canonical:0001-com-ubuntu-server-focal:20_04-lts-gen2:latest \
    --os-disk-size-gb 128 \
    --storage-sku Premium_LRS \
    --generate-ssh-keys \
    --public-ip-sku Standard \
    --query publicIpAddress \
    -o tsv \
    --custom-data cluster-$Store.sh)

  echo -e "$Store\t$IP" >> ips

  _information "VM Created . IP: $IP"

  _information "Deleting default SSH rule"
  az network nsg rule delete -g $Store --nsg-name ${Store}NSG -o table --name default-allow-ssh

  # For more security, replace --source-address-prefixes * with your IP or CIDR

  _information "Creating SSH rule on port 2222"
  az network nsg rule create \
  -g $Store \
  --nsg-name ${Store}NSG \
  -n SSH-http \
  --description "SSH http https" \
  --destination-port-ranges 2222 80 443 \
  --protocol tcp \
  --access allow \
  --priority 1202


  if [ ! -z "$DnsZone" ]; then
    _information "Configuring DNS"
    _information "DnsZone: $DnsZone Name: $Store"
    upsert_dns_record
    
    if [ "$UseSSL" == "true" ]; then
      _information "Copying SSL Cert Files to VM"
      copy_cert_files_to_vm
      wait_for_cloud_init_completion "$IP"
      _information "Creating SSL Cert secret"
      create_cert_secret "$IP"
    fi
  fi

  _success "VM Creation Complete! $Store  $IP"
  _success "akdc ssh $Store"
  
  rm -f cluster-$Store.sh
}

validate_input() {
  if [ "$ShowHelp" == "true" ]; then
    usage
  fi

  if [ -z $AKDC_PAT ]; then
    usage
  fi

  if [ $_argLength -lt 4 ]; then
    _error "Missing input parameters"
    usage
  fi

  if [ ! -z "$DnsZone" ]; then
    if [ -z "$CertCRTPath" ]; then
      _error "SSL Cert .crt path is a required argument when using --zone"
      usage
    fi
    if [ -z "$CertKeyPath" ]; then
      _error "SSL Cert .key path is a required argument when using --zone"
      usage
    fi    
  fi  
}

create_cert_secret() {
  local ip=$1
  ssh -p 2222 -o "StrictHostKeyChecking=no" akdc@$ip 'kubectl create secret generic ssl-cert --from-file=tls.crt=/home/akdc/zone.crt --from-file=tls.key=/home/akdc/zone.key && rm /home/akdc/zone.key && rm /home/akdc/zone.crt'
}

wait_for_cloud_init_completion() {
  local ip=$1
  _information "Waiting for cloud init to complete."

  status=$(ssh -p 2222 -o "StrictHostKeyChecking=no" akdc@$ip 'cloud-init status')
  _debug "cloud init status:$status."
  while [ "$status" != "status: done" ]; do
    sleep 10
    status=$(ssh -p 2222 -o "StrictHostKeyChecking=no" akdc@$ip 'cloud-init status')
    _debug "cloud init status:$status."
  done
}

tool_installed() {
  tool=$1
    local jqStatusMessage
    if [ -x "$(command -v $tool)" ]; then
        jqStatusMessage="(installed - you're good to go!)"
    else
        jqStatusMessage="\e[31m(not installed)\e[0m"
    fi   
    echo $jqStatusMessage   
}

usage() {
    local jqStatusMessage=$(tool_installed "jq")    
    local sedStatusMessage=$(tool_installed "sed")
    
    _helpText="  Usage: $me Region State City Store-Number AzureRegion

  Positional Arguments
          Region        Required
           State        Required
            City        Required
    Store-Number        Required
     AzureRegion        Optional. Defaults to centralus if not supplied

  Optional Named Arguments
    -z | --zone       <dns domain>          The dns domain to be configured on this VM. If omitted DNS will not configured. This script assumes an Azure zone with the domain as it's name.
    -s | --ssl                              Configure an SSL cert for this domain. Cert CRT and Key must be provied
    -l | --location   <Azure Location>      An alternative to specifying the AzureRegion via a named argument instead of the 5th postional argument.
    -c | --crt        <path to .crt file>   File Path to SSL Cert .crt file. Only required if --ssl is used
    -k | --key        <path to .key file>   File Path to SSL Cert .key file. Only required if --ssl is used. This must be a decrypted key!
  dependencies:
    -jq $jqStatusMessage
    -sed $sedStatusMessage
         "
                
    _information "$_helpText" 1>&2
    exit 1
}

banner() {
  local version=$1
  cat << "EOF"

 █████  ██   ██ ██████   ██████ 
██   ██ ██  ██  ██   ██ ██      
███████ █████   ██   ██ ██      
██   ██ ██  ██  ██   ██ ██      
██   ██ ██   ██ ██████   ██████ 

EOF

  echo "  version: $(cat ./version.txt)"
  echo "  k3d in VMs "   
  echo ""                                                      
}

copy_cert_files_to_vm() {  
  scp -o "StrictHostKeyChecking=no" -P 2222 $CertCRTPath akdc@$IP:~/zone.crt
  scp -o "StrictHostKeyChecking=no" -P 2222 $CertKeyPath akdc@$IP:~/zone.key
}

upsert_dns_record() {
   OLD_IP=$(az network dns record-set a list \
  --query "[?name=='$Store'].{IP:aRecords}" \
  --resource-group "$DnsZoneResourceGroup" \
  --zone-name "$DnsZone"  -o json | jq -r '.[].IP[].ipv4Address')

  az network dns record-set a add-record \
    -g $DnsZoneResourceGroup \
    -z $DnsZone \
    -n $Store \
    -a $IP \
    --ttl $DnsRecordTTL

  if [[ ! -z "$OLD_IP" ]]; then
  _information "Removing old IP: $OLD_IP from record $Store"
  az network dns record-set a remove-record \
  -g $DnsZoneResourceGroup \
  -z  $DnsZone \
  -n $Store \
  -a $OLD_IP
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

_information "  @Store Context"
echo "             Region: $Region"
echo "              State: $State"
echo "               City: $City"
echo "             Number: $Number"
echo "           Location: $Location"
echo "           District: $District"
echo "              Store: $Store"
echo "            DnsZone: $DnsZone"
echo " DNS Resource Group: $DnsZoneResourceGroup"
echo "            DNS TTL: $DnsRecordTTL"
echo "               FQDN: $FQDN"
}

main
