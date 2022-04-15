#!/bin/bash

### run as su - will not work with sudo

set -e

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-config start" >> /home/akdc/status

# configure git CLI
git config --system user.name autogitops
git config --system user.email autogitops@outlook.com
git config --system core.whitespace blank-at-eol,blank-at-eof,space-before-tab
git config --system pull.rebase false
git config --system init.defaultbranch main
git config --system fetch.prune true
git config --system core.pager more

# make some directories we will need
mkdir -p .ssh
mkdir -p .kube
mkdir -p bin
mkdir -p .local/bin
mkdir -p .k9s
mkdir -p .oh-my-bash/completions
mkdir -p /root/.kube

chown -R "${AKDC_ME}:${AKDC_ME}" "/home/${AKDC_ME}"

# make some system dirs
mkdir -p /etc/docker
mkdir -p /prometheus && chown -R 65534:65534 /prometheus
mkdir -p /grafana
chown -R 472:0 /grafana

# create / add to groups
groupadd docker
usermod -aG sudo "${AKDC_ME}"
usermod -aG admin "${AKDC_ME}"
usermod -aG docker "${AKDC_ME}"
gpasswd -a "${AKDC_ME}" sudo

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-config complete" >> /home/akdc/status
