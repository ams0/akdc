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
if [ ! -f "$HOME"/cli/vm/setup/akdc-install.sh ]
then
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-install.sh not found" >> "/home/${AKDC_ME}/status"
  echo "akdc-install.sh not found"
  exit 1
fi

# can't continue without akdc-dns.sh
if [ ! -f "$HOME"/cli/vm/setup/akdc-dns.sh ]
then
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-dns.sh not found" >> "/home/${AKDC_ME}/status"
  echo "akdc-dns.sh not found"
  exit 1
fi

# can't continue without k8s-setup.sh
if [ ! -f "$HOME"/cli/vm/setup/k8s-setup.sh ]
then
  echo "$(date +'%Y-%m-%d %H:%M:%S')  k8s-setup.sh not found" >> "/home/${AKDC_ME}/status"
  echo "k8s-setup.sh not found"
  exit 1
fi

# can't continue without flux-setup.sh
if [ ! -f "$HOME"/cli/vm/setup/flux-setup.sh ]
then
  echo "$(date +'%Y-%m-%d %H:%M:%S')  flux-setup.sh not found" >> "/home/${AKDC_ME}/status"
  echo "flux-setup.sh not found"
  exit 1
fi

set -e

# run setup scripts
"$HOME"/cli/vm/setup/akdc-install.sh
"$HOME"/cli/vm/setup/akdc-dns.sh

# don't run the rest of setup in debug mode
if [ "$AKDC_DEBUG" = "true" ]
then
  echo "$(date +'%Y-%m-%d %H:%M:%S')  debug mode" >> status
  echo "debug mode"
  exit 0
fi

# run akdc-pre-k8s.sh
if [ -f "$HOME/cli/vm/setup/akdc-pre-k8s.sh" ]
then
  # run as AKDC_ME
  "$HOME"/cli/vm/setup/akdc-pre-k8s.sh
else
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre-k8s.sh not found" >> "/home/${AKDC_ME}/status"
fi

# run k8s-setup
"$HOME"/cli/vm/setup/k8s-setup.sh

# run akdc-pre-flux.sh
if [ -f "$HOME"/cli/vm/setup/akdc-pre-flux.sh ]
then
  "$HOME"/cli/vm/setup/akdc-pre-flux.sh
else
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre-flux.sh not found" >> "/home/${AKDC_ME}/status"
fi

# setup flux
"$HOME"/cli/vm/setup/flux-setup.sh

# run akdc-pre-arc.sh
if [ -f "$HOME"/cli/vm/setup/akdc-pre-arc.sh ]
then
  "$HOME"/cli/vm/setup/akdc-pre-arc.sh
else
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre-arc.sh not found" >> "/home/${AKDC_ME}/status"
fi

# setup azure arc
if [ -f "$HOME"/cli/vm/setup/arc-setup.sh ]
then
  "$HOME"/cli/vm/setup/arc-setup.sh
else
  echo "$(date +'%Y-%m-%d %H:%M:%S')  arc-setup.sh not found" >> "/home/${AKDC_ME}/status"
fi

# run akdc-private-repos.sh
if [ -f "$HOME"/cli/vm/setup/akdc-private-repos.sh ]
then
  "$HOME"/cli/vm/setup/akdc-private-repos.sh
else
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-private-repos.sh not found" >> "/home/${AKDC_ME}/status"
fi

# run akdc-post.sh
if [ -f "$HOME"/cli/vm/setup/akdc-post.sh ]
then
  "$HOME"/cli/vm/setup/akdc-post.sh
else
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-post.sh not found" >> "/home/${AKDC_ME}/status"
fi

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-setup complete" >> "/home/${AKDC_ME}/status"
echo "complete" >> "/home/${AKDC_ME}/status"
