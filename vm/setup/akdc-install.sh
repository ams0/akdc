#!/bin/bash

# this is the main VM setup script

# run as su - will not work with sudo

set -e

echo "$(date +'%Y-%m-%d %H:%M:%S')  installing libs" >> status
apt-get install -y net-tools software-properties-common libssl-dev libffi-dev python-dev build-essential lsb-release gnupg-agent

echo "$(date +'%Y-%m-%d %H:%M:%S')  installing utils" >> status
apt-get install -y curl git wget nano jq zip unzip httpie
apt-get install -y dnsutils coreutils gnupg2 make bash-completion gettext iputils-ping

echo "$(date +'%Y-%m-%d %H:%M:%S')  adding package sources" >> status

# add Docker repo
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key --keyring /etc/apt/trusted.gpg.d/docker.gpg add -
add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"

# add kubenetes repo
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" > /etc/apt/sources.list.d/kubernetes.list

apt-get update

echo "$(date +'%Y-%m-%d %H:%M:%S')  installing docker" >> status
apt-get install -y docker-ce docker-ce-cli

echo "$(date +'%Y-%m-%d %H:%M:%S')  installing kubectl" >> status
apt-get install -y kubectl

# Install istio CLI
echo "$(date +'%Y-%m-%d %H:%M:%S')  installing istioctl" >> status
echo "Installing istioctl"
curl -sL https://istio.io/downloadIstioctl | sh -
mv ~/.istioctl/bin/istioctl /usr/local/bin

# kubectl auto complete
#kubectl completion bash > /etc/bash_completion.d/kubectl

echo "$(date +'%Y-%m-%d %H:%M:%S')  installing k3d" >> status
wget -q -O - https://raw.githubusercontent.com/rancher/k3d/main/install.sh | TAG=v4.4.8 bash

echo "$(date +'%Y-%m-%d %H:%M:%S')  installing flux" >> status
curl -s https://fluxcd.io/install.sh | bash
#flux completion bash > "/home/${AKDC_ME}/.oh-my-bash/completions/flux.sh"

echo "$(date +'%Y-%m-%d %H:%M:%S')  installing k9s" >> status
VERSION=$(curl -i https://github.com/derailed/k9s/releases/latest | grep "location: https://github.com/" | rev | cut -f 1 -d / | rev | sed 's/\r//')
wget "https://github.com/derailed/k9s/releases/download/${VERSION}/k9s_Linux_x86_64.tar.gz"
tar -zxvf k9s_Linux_x86_64.tar.gz -C /usr/local/bin
rm -f k9s_Linux_x86_64.tar.gz

# upgrade Ubuntu
echo "$(date +'%Y-%m-%d %H:%M:%S')  upgrading" >> status
apt-get update
apt-get upgrade -y
apt-get autoremove -y

echo "$(date +'%Y-%m-%d %H:%M:%S')  creating registry" >> status
# create local registry
chown -R "${AKDC_ME}:${AKDC_ME}" "/home/${AKDC_ME}"
docker network create k3d
k3d registry create registry.localhost --port 5500
docker network connect k3d k3d-registry.localhost
