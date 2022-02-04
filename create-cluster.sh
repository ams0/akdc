#!/bin/bash

if [ -z $AKDC_PAT ]; then
  echo "Please export AKDC_PAT env variables"
  exit 1
fi

if [ $# -lt 4 ]; then
  echo "Usage: $0 Region State City Store-Number [AzureRegion:centralus]"
  exit 1
fi

Region=${1}
State=${2}
City=${3}
Number=${4}
Location=${5:-centralus}
District=${Region}-${State}-${City}
Store=${District}-${Number}

# create the RG
az group create -l $Location -n $Store

# create the install script from the template

# replace the host, pat, district and region
rm -f cluster-$Store.sh
sed "s/{{pat}}/$AKDC_PAT/g" ./akdc.templ | \
    sed "s/{{store}}/$Store/g" | \
    sed "s/{{district}}/$District/g" | \
    sed "s/{{region}}/$Region/g" \
    > cluster-$Store.sh

# create the VM
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

# delete SSH rule
az network nsg rule delete -g $Store --nsg-name ${Store}NSG -o table --name default-allow-ssh

# For more security, replace --source-address-prefixes * with your IP or CIDR

# create SSH rule on port 2222
az network nsg rule create \
-g $Store \
--nsg-name ${Store}NSG \
-n SSH-http \
--description "SSH http https" \
--destination-port-ranges 2222 80 443 \
--protocol tcp \
--access allow \
--priority 1202

echo "$Store  $IP"
echo -e "$Store\t$IP" >> ips
