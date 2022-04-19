#!/bin/bash

# this script installs most of the components

set -e

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-install start" >> "/home/${AKDC_ME}/status"

echo "$(date +'%Y-%m-%d %H:%M:%S')  installing libs" >> "/home/${AKDC_ME}/status"
sudo apt-get install -y net-tools software-properties-common libssl-dev libffi-dev python-dev build-essential lsb-release gnupg-agent

echo "$(date +'%Y-%m-%d %H:%M:%S')  installing utils" >> "/home/${AKDC_ME}/status"
sudo apt-get install -y curl git wget nano jq zip unzip httpie
sudo apt-get install -y dnsutils coreutils gnupg2 make bash-completion gettext iputils-ping

echo "$(date +'%Y-%m-%d %H:%M:%S')  adding package sources" >> "/home/${AKDC_ME}/status"

# add Docker repo
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key --keyring /etc/apt/trusted.gpg.d/docker.gpg add -
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"

# add kubenetes repo
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list

sudo apt-get update

echo "$(date +'%Y-%m-%d %H:%M:%S')  installing docker" >> "/home/${AKDC_ME}/status"
sudo apt-get install -y docker-ce docker-ce-cli

echo "$(date +'%Y-%m-%d %H:%M:%S')  installing kubectl" >> "/home/${AKDC_ME}/status"
sudo apt-get install -y kubectl

# Install istio CLI
echo "$(date +'%Y-%m-%d %H:%M:%S')  installing istioctl" >> "/home/${AKDC_ME}/status"
echo "Installing istioctl"
curl -sL https://istio.io/downloadIstioctl | bash -
sudo mv "/home/${AKDC_ME}/.istioctl/bin/istioctl" /usr/local/bin

echo "$(date +'%Y-%m-%d %H:%M:%S')  installing k3d" >> "/home/${AKDC_ME}/status"
wget -q -O - https://raw.githubusercontent.com/rancher/k3d/main/install.sh | sudo TAG=v4.4.8 bash

echo "$(date +'%Y-%m-%d %H:%M:%S')  installing flux" >> "/home/${AKDC_ME}/status"
curl -s https://fluxcd.io/install.sh | sudo bash

echo "$(date +'%Y-%m-%d %H:%M:%S')  installing k9s" >> "/home/${AKDC_ME}/status"
VERSION=$(curl -i https://github.com/derailed/k9s/releases/latest | grep "location: https://github.com/" | rev | cut -f 1 -d / | rev | sed 's/\r//')
wget "https://github.com/derailed/k9s/releases/download/${VERSION}/k9s_Linux_x86_64.tar.gz"
sudo tar -zxvf k9s_Linux_x86_64.tar.gz -C /usr/local/bin
rm -f k9s_Linux_x86_64.tar.gz

### new code
sudo mkdir -p /etc/caddy

# add caddy sources
sudo apt-get install -y debian-keyring debian-archive-keyring apt-transport-https
curl -1sLf https://dl.cloudsmith.io/public/caddy/stable/gpg.key | sudo tee /etc/apt/trusted.gpg.d/caddy-stable.asc
curl -1sLf https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt | sudo tee /etc/apt/sources.list.d/caddy-stable.list

# add dotnet repo
echo "deb [arch=amd64] https://packages.microsoft.com/repos/microsoft-ubuntu-$(lsb_release -cs)-prod $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/dotnetdev.list

# upgrade Ubuntu
echo "$(date +'%Y-%m-%d %H:%M:%S')  upgrading" >> "/home/${AKDC_ME}/status"
sudo apt-get update
sudo apt-get upgrade -y
sudo apt-get autoremove -y

cat << EOF | sudo tee /etc/caddy/Caddyfile
${AKDC_FQDN} {
  redir /heartbeat /heartbeat/
  redir /webv /webv/
  redir /grafana /grafana/
  redir /prometheus /prometheus/
  reverse_proxy 127.0.0.1:30080
}

${AKDC_FQDN}/heartbeat/* {
	reverse_proxy 127.0.0.1:30082
}

${AKDC_FQDN}/grafana/* {
        uri strip_prefix /grafana
        reverse_proxy 127.0.0.1:32000
}

${AKDC_FQDN}/prometheus/* {
        reverse_proxy 127.0.0.1:30000
}

${AKDC_FQDN}/webv/* {
        uri strip_prefix /webv
        reverse_proxy 127.0.0.1:30088
}
EOF

echo "installing dotnet" >> "/home/${AKDC_ME}/status"
sudo apt-get install -y dotnet-sdk-6.0

dotnet tool install -g webvalidate

# install caddy
sudo apt-get install -y caddy

### end new code

sudo chown -R "${AKDC_ME}:${AKDC_ME}" "/home/${AKDC_ME}"

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-install complete" >> "/home/${AKDC_ME}/status"
