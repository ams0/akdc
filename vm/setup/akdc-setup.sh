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

set -e

# run setup scripts
./cli/vm/setup/akdc-install.sh
./cli/vm/setup/akdc-dns.sh

# run akdc-post.sh
if [ -f ./cli/vm/setup/akdc-post.sh ]
then
  ./cli/vm/setup/akdc-post.sh
else
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-post.sh not found" >> "/home/${AKDC_ME}/status"
fi

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-setup complete" >> "/home/${AKDC_ME}/status"
echo "complete" >> "/home/${AKDC_ME}/status"
