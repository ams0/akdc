#!/bin/bash

### to run manually
# cd /home/akdc
# cli/vm/setup/akdc-setup.sh

# this is the main VM setup script

# env variables defined in /etc/bash.bashrc
    # AKDC_ARC_ENABLED
    # AKDC_BRANCH
    # AKDC_CLUSTER
    # AKDC_DEBUG
    # AKDC_DNS_RG
    # AKDC_FQDN
    # AKDC_ME
    # AKDC_REPO
    # AKDC_RESOURCE_GROUP
    # AKDC_ZONE

# change to this directory
# cd "$(dirname "${BASH_SOURCE[0]}")" || exit

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-setup start" >> "/home/${AKDC_ME}/status"

# can't continue without akdc-install.sh
if [ ! -f ./cli/vm/setup/akdc-install.sh ]
then
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-install.sh not found" >> "/home/${AKDC_ME}/status"
  echo "akdc-install.sh not found"
  exit 1
fi

# can't continue without akdc-dns.sh
if [ ! -f ./cli/vm/setup/akdc-dns.sh ]
then
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-dns.sh not found" >> "/home/${AKDC_ME}/status"
  echo "akdc-dns.sh not found"
  exit 1
fi

# can't continue without k3d-setup.sh
if [ ! -f ./cli/vm/setup/k3d-setup.sh ]
then
  echo "$(date +'%Y-%m-%d %H:%M:%S')  k3d-setup.sh not found" >> "/home/${AKDC_ME}/status"
  echo "k3d-setup.sh not found"
  exit 1
fi

# can't continue without flux-setup.sh
if [ ! -f ./cli/vm/setup/flux-setup.sh ]
then
  echo "$(date +'%Y-%m-%d %H:%M:%S')  flux-setup.sh not found" >> "/home/${AKDC_ME}/status"
  echo "flux-setup.sh not found"
  exit 1
fi

set -e

# run setup scripts
./cli/vm/setup/akdc-install.sh
./cli/vm/setup/akdc-dns.sh

# set env vars
export AKDC_RESOURCE_GROUP=factory-fleet
export AKDC_STORAGE_NAME=factoryfleetstorage
export AKDC_VOLUME=uploadvolume
export AKDC_SP_ID=$(az keyvault secret show --vault-name kv-tld  --query 'value' -o tsv -n akdc-sp-id)
export AKDC_SP_KEY=$(az keyvault secret show --vault-name kv-tld  --query 'value' -o tsv -n akdc-sp-key)
export AKDC_TENANT=$(az account show --query id -o tsv)
export REPO_BASE=/workspaces/edge-gitops
export AKDC_SUBSCRIPTION=$(az account show --query id -o tsv)
export AKDC_STORAGE_KEY=$(az storage account keys list --resource-group "$AKDC_RESOURCE_GROUP" --account-name "$AKDC_STORAGE_NAME" --query "[0].value" -o tsv)
export AKDC_STORAGE_CONNECTION=$(az storage account show-connection-string -n "$AKDC_STORAGE_NAME" -g "$AKDC_RESOURCE_GROUP" -o tsv)

# run akdc-pre-k3d.sh
if [ -f ./cli/vm/setup/akdc-pre-k3d.sh ]
then
  # run as AKDC_ME
  ./cli/vm/setup/akdc-pre-k3d.sh
else
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre-k3d.sh not found" >> "/home/${AKDC_ME}/status"
fi

# run k3d-setup
./cli/vm/setup/k3d-setup.sh

# run akdc-pre-flux.sh
if [ -f ./cli/vm/setup/akdc-pre-flux.sh ]
then
  ./cli/vm/setup/akdc-pre-flux.sh
else
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre-flux.sh not found" >> "/home/${AKDC_ME}/status"
fi

# setup flux
./cli/vm/setup/flux-setup.sh

# run akdc-pre-arc.sh
if [ -f ./cli/vm/setup/akdc-pre-arc.sh ]
then
  ./cli/vm/setup/akdc-pre-arc.sh
else
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre-arc.sh not found" >> "/home/${AKDC_ME}/status"
fi

# setup azure arc
if [ -f ./cli/vm/setup/arc-setup.sh ]
then
  ./cli/vm/setup/arc-setup.sh
else
  echo "$(date +'%Y-%m-%d %H:%M:%S')  arc-setup.sh not found" >> "/home/${AKDC_ME}/status"
fi

# run akdc-private-repos.sh
if [ -f ./cli/vm/setup/akdc-private-repos.sh ]
then
  ./cli/vm/setup/akdc-private-repos.sh
else
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-private-repos.sh not found" >> "/home/${AKDC_ME}/status"
fi

# run akdc-post.sh
if [ -f ./cli/vm/setup/akdc-post.sh ]
then
  ./cli/vm/setup/akdc-post.sh
else
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-post.sh not found" >> "/home/${AKDC_ME}/status"
fi

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-setup complete" >> "/home/${AKDC_ME}/status"
echo "complete" >> "/home/${AKDC_ME}/status"
