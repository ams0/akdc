#!/bin/bash

# this script installs most of the components

set -e

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-install start" >> "/home/${AKDC_ME}/status"

echo "$(date +'%Y-%m-%d %H:%M:%S')  installing libs" >> "/home/${AKDC_ME}/status"
sudo apt-get install -y net-tools software-properties-common libssl-dev libffi-dev python-dev build-essential lsb-release gnupg-agent

echo "$(date +'%Y-%m-%d %H:%M:%S')  installing utils" >> "/home/${AKDC_ME}/status"
sudo apt-get install -y curl git wget nano jq zip unzip httpie
sudo apt-get install -y dnsutils coreutils gnupg2 make bash-completion gettext iputils-ping

# add Docker source
echo "$(date +'%Y-%m-%d %H:%M:%S')  adding docker source" >> "/home/${AKDC_ME}/status"
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key --keyring /etc/apt/trusted.gpg.d/docker.gpg add -
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"

# add kubenetes source
echo "$(date +'%Y-%m-%d %H:%M:%S')  adding kubernetes source" >> "/home/${AKDC_ME}/status"
curl -fsSL https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list

echo "$(date +'%Y-%m-%d %H:%M:%S')  updating sources" >> "/home/${AKDC_ME}/status"

# this is failing on large fleets - add one retry

set +e

if ! sudo apt-get update
then
    echo "$(date +'%Y-%m-%d %H:%M:%S')  updating sources (retry)" >> "/home/${AKDC_ME}/status"
    sleep 30
    set -e
    sudo apt-get update
fi

set -e

echo "$(date +'%Y-%m-%d %H:%M:%S')  installing docker" >> "/home/${AKDC_ME}/status"
sudo apt-get install -y docker-ce docker-ce-cli

echo "$(date +'%Y-%m-%d %H:%M:%S')  installing kubectl" >> "/home/${AKDC_ME}/status"
sudo apt-get install -y kubectl

echo "$(date +'%Y-%m-%d %H:%M:%S')  installing flux" >> "/home/${AKDC_ME}/status"
curl -s https://fluxcd.io/install.sh | sudo bash

echo "$(date +'%Y-%m-%d %H:%M:%S')  installing k9s" >> "/home/${AKDC_ME}/status"
VERSION=$(curl -i https://github.com/derailed/k9s/releases/latest | grep "location: https://github.com/" | rev | cut -f 1 -d / | rev | sed 's/\r//')
wget "https://github.com/derailed/k9s/releases/download/${VERSION}/k9s_Linux_x86_64.tar.gz"
sudo tar -zxvf k9s_Linux_x86_64.tar.gz -C /usr/local/bin
rm -f k9s_Linux_x86_64.tar.gz

# upgrade Ubuntu
echo "$(date +'%Y-%m-%d %H:%M:%S')  upgrading" >> "/home/${AKDC_ME}/status"
sudo apt-get update
sudo apt-get upgrade -y
sudo apt-get autoremove -y

sudo chown -R "${AKDC_ME}:${AKDC_ME}" "/home/${AKDC_ME}"

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-install complete" >> "/home/${AKDC_ME}/status"
