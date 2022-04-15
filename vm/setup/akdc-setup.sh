#!/bin/bash

### to run manually
# cd /home/akdc
# sudo su
# cli/vm/setup/akdc-setup.sh

### will not work via sudo ###

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

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-setup start" >> /home/akdc/status

set -e

# can't continue without akdc-config.sh
if [ ! -f ./cli/vm/setup/akdc-config.sh ]
then
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-config.sh not found" >> /home/akdc/status
  echo "akdc-config.sh not found"
  exit 1
fi

# can't continue without akdc-install.sh
if [ ! -f ./cli/vm/setup/akdc-install.sh ]
then
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-install.sh not found" >> /home/akdc/status
  echo "akdc-install.sh not found"
  exit 1
fi

# can't continue without akdc-dns.sh
if [ ! -f ./cli/vm/setup/akdc-dns.sh ]
then
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-dns.sh not found" >> /home/akdc/status
  echo "akdc-dns.sh not found"
  exit 1
fi

# can't continue without k3d-setup.sh
if [ ! -f ./cli/vm/setup/k3d-setup.sh ]
then
  echo "$(date +'%Y-%m-%d %H:%M:%S')  k3d-setup.sh not found" >> /home/akdc/status
  echo "k3d-setup.sh not found"
  exit 1
fi

# can't continue without flux-setup.sh
if [ ! -f ./cli/vm/setup/flux-setup.sh ]
then
  echo "$(date +'%Y-%m-%d %H:%M:%S')  flux-setup.sh not found" >> /home/akdc/status
  echo "flux-setup.sh not found"
  exit 1
fi

# run setup scripts
./cli/vm/setup/akdc-config.sh
./cli/vm/setup/akdc-install.sh
./cli/vm/setup/akdc-dns.sh

# run akdc-pre-k3d.sh
if [ -f ./cli/vm/setup/akdc-pre-k3d.sh ]
then
  # run as AKDC_ME
  sudo -HEu akdc ./cli/vm/setup/akdc-pre-k3d.sh || exit 1
else
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre-k3d.sh not found" >> /home/akdc/status
fi

# run k3d-setup
sudo -HEu akdc ./cli/vm/setup/k3d-setup.sh || exit 1

# run akdc-pre-flux.sh
if [ -f ./cli/vm/setup/akdc-pre-flux.sh ]
then
  sudo -HEu akdc ./cli/vm/setup/akdc-pre-flux.sh || exit 1
else
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre-flux.sh not found" >> /home/akdc/status
fi

# setup flux
sudo -HEu akdc ./cli/vm/setup/flux-setup.sh || exit 1

# run akdc-pre-arc.sh
if [ -f ./cli/vm/setup/akdc-pre-arc.sh ]
then
  sudo -HEu akdc ./cli/vm/setup/akdc-pre-arc.sh || exit 1
else
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre-arc.sh not found" >> /home/akdc/status
fi

# configure azure arc
if [ -f ./cli/vm/setup/arc-setup.sh ]
then
  sudo -HEu akdc ./cli/vm/setup/arc-setup.sh || exit 1
else
  echo "$(date +'%Y-%m-%d %H:%M:%S')  arc-setup.sh not found" >> /home/akdc/status
fi

# run akdc-private-repos.sh
if [ -f ./cli/vm/setup/akdc-private-repos.sh ]
then
  sudo -HEu akdc ./cli/vm/setup/akdc-private-repos.sh || exit 1
else
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-private-repos.sh not found" >> /home/akdc/status
fi

# run akdc-post.sh
if [ -f ./cli/vm/setup/akdc-post.sh ]
then
  sudo -HEu akdc ./cli/vm/setup/akdc-post.sh || exit 1
else
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-post.sh not found" >> /home/akdc/status
fi

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-setup complete" >> /home/akdc/status
echo "complete" >> /home/akdc/status
