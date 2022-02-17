#!/bin/sh

####### do not change these values #######
export ME=akdc
export FQDN="{{fqdn}}"
###################################

export DEBIAN_FRONTEND=noninteractive
export HOME=/root

### Needed for Digital Ocean
useradd -m -s /bin/bash ${ME}
mkdir -p /home/${ME}/.ssh
cp /root/.ssh/authorized_keys /home/${ME}/.ssh
### end DO

cp /usr/share/zoneinfo/America/Chicago /etc/localtime

cd /home/${ME} || exit

echo "$(date)  starting" >> status

echo "${ME} ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers.d/akdc

# upgrade sshd security
{
  echo ""
  echo "ClientAliveInterval 120"
  echo "Port 2222"
  echo "Ciphers chacha20-poly1305@openssh.com,aes256-gcm@openssh.com,aes128-gcm@openssh.com,aes256-ctr,aes192-ctr,aes128-ctr"
} >> /etc/ssh/sshd_config

# restart sshd
systemctl restart sshd

# make some directories we will need
mkdir -p .ssh
mkdir -p .kube
mkdir -p bin
mkdir -p .local/bin
mkdir -p .k9s
mkdir -p /root/.kube

# configure git CLI
git config --global user.name autogitops
git config --global user.email autogitops@outlook.com
git config --system core.whitespace blank-at-eol,blank-at-eof,space-before-tab
git config --system pull.rebase false
git config --system init.defaultbranch main
git config --system fetch.prune true
git config --system core.pager more

# oh my bash
git clone --depth=1 https://github.com/ohmybash/oh-my-bash.git .oh-my-bash
cp .oh-my-bash/templates/bashrc.osh-template .bashrc

# add to .bashrc
# shellcheck disable=2016
{
  echo ""
  echo "alias k='kubectl'"
  echo "alias kga='kubectl get all'"
  echo "alias kgaa='kubectl get all --all-namespaces'"
  echo "alias kaf='kubectl apply -f'"
  echo "alias kdelf='kubectl delete -f'"
  echo "alias kl='kubectl logs'"
  echo "alias kccc='kubectl config current-context'"
  echo "alias kcgc='kubectl config get-contexts'"
  echo "alias kj='kubectl exec -it jumpbox -- bash -l'"
  echo "alias kje='kubectl exec -it jumpbox -- '"
  echo "alias sync='flux reconcile source git gitops && kubectl get pods -A'"

  echo ""
  echo "alias ipconfig='ip -4 a show eth0 | grep inet | sed \"s/inet//g\" | sed \"s/ //g\" | cut -d / -f 1'"

  echo ""
  echo 'export PIP=$(ipconfig | tail -n 1)'
  echo 'export PATH="$PATH:/usr/local/istio/bin:$HOME/.dotnet/tools:$HOME/go/bin"'
  echo "export FQDN=$FQDN"

  echo ""
  echo 'complete -F __start_kubectl k'
  echo 'complete -F __start_kubectl kl'
} >> .bashrc

chown -R ${ME}:${ME} /home/${ME}

# make some system dirs
mkdir -p /etc/docker
mkdir -p /prometheus && chown -R 65534:65534 /prometheus
mkdir -p /grafana
# cp /workspaces/.cnp-labs/cluster-admin/deploy/grafanadata/grafana.db /grafana
chown -R 472:0 /grafana

# create / add to groups
groupadd docker
usermod -aG sudo ${ME}
usermod -aG admin ${ME}
usermod -aG docker ${ME}
gpasswd -a ${ME} sudo

echo "$(date)  installing base" >> status
apt-get update
apt-get install -y apt-utils dialog apt-transport-https ca-certificates net-tools

echo "$(date)  installing libs" >> status
apt-get install -y software-properties-common libssl-dev libffi-dev python-dev build-essential lsb-release gnupg-agent

echo "$(date)  installing utils" >> status
apt-get install -y curl git wget nano jq zip unzip httpie
apt-get install -y dnsutils coreutils gnupg2 make bash-completion gettext iputils-ping

echo "$(date)  adding package sources" >> status

# add Docker repo
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key --keyring /etc/apt/trusted.gpg.d/docker.gpg add -
add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"

# add kubenetes repo
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" > /etc/apt/sources.list.d/kubernetes.list

apt-get update

echo "$(date)  installing docker" >> status
apt-get install -y docker-ce docker-ce-cli

echo "$(date)  installing kubectl" >> status
apt-get install -y kubectl

# kubectl auto complete
kubectl completion bash > /etc/bash_completion.d/kubectl

echo "$(date)  installing k3d" >> status
wget -q -O - https://raw.githubusercontent.com/rancher/k3d/main/install.sh | TAG=v4.4.8 bash

echo "$(date)  creating registry" >> status
# create local registry
chown -R ${ME}:${ME} /home/${ME}
docker network create k3d
k3d registry create registry.localhost --port 5500
docker network connect k3d k3d-registry.localhost

echo "$(date)  creating k3d cluster" >> status

cat << EOF > k3d.yaml
apiVersion: k3d.io/v1alpha2
kind: Simple
servers: 1
network: k3d
kubeAPI:
  hostIP: "0.0.0.0"
  hostPort: "6443"
volumes:
  - volume: /prometheus:/prometheus
    nodeFilters:
      - server[0]
  - volume: /grafana:/grafana
    nodeFilters:
      - server[0]
ports:
  - port: 443:443
    nodeFilters:
      - loadbalancer
  - port: 80:80 
    nodeFilters:
      - loadbalancer
  - port: 30000:30000
    nodeFilters:
      - server[0]
  - port: 30080:30080
    nodeFilters:
      - server[0]
  - port: 30088:30088
    nodeFilters:
      - server[0]
  - port: 32000:32000
    nodeFilters:
      - server[0]

options:
  k3d:
    wait: true
    timeout: "60s"
  k3s:
    extraServerArgs:
      - --tls-san=127.0.0.1
    extraAgentArgs: []
  kubeconfig:
    updateDefaultKubeconfig: true
    switchCurrentContext: true
EOF

# change ownership
chown -R ${ME}:${ME} /home/${ME}

k3d cluster create \
--registry-use k3d-registry.localhost:5500 \
--config k3d.yaml 

# copy kube config file
cp /root/.kube/config /home/${ME}/.kube/config
chown -R ${ME}:${ME} /home/${ME}

echo "$(date)  installing flux" >> status
curl -s https://fluxcd.io/install.sh | bash
flux completion bash > /etc/bash_completion.d/flux

#echo "$(date)  installing istio" >> status
#curl -L https://istio.io/downloadIstio | sh -
#mv istio* istio
#chmod -R 755 istio
#cp istio/tools/istioctl.bash /etc/bash_completion.d
#chown -R ${ME}:${ME} /home/${ME}
#mv istio /usr/local

echo "$(date)  installing k9s" >> status
VERSION=$(curl -i https://github.com/derailed/k9s/releases/latest | grep "location: https://github.com/" | rev | cut -f 1 -d / | rev | sed 's/\r//')
wget "https://github.com/derailed/k9s/releases/download/${VERSION}/k9s_Linux_x86_64.tar.gz"
tar -zxvf k9s_Linux_x86_64.tar.gz -C /usr/local/bin
rm -f k9s_Linux_x86_64.tar.gz

# upgrade Ubuntu
echo "$(date)  upgrading" >> status
apt-get update
apt-get upgrade -y
apt-get autoremove -y

echo "$(date)  waiting for cluster to start" >> status
kubectl wait node --for condition=ready --all --timeout=30s

echo "$(date)  installing dapr" >> status
wget -q https://raw.githubusercontent.com/dapr/cli/master/install/install.sh -O - | /bin/bash
dapr init -k --enable-mtls=false --wait

echo "$(date)  installing radius" >> status
wget -q "https://get.radapp.dev/tools/rad/install.sh" -O - | /bin/bash
rad env init kubernetes -n radius-system

echo "$(date)  waiting for pods to start" >> status
kubectl wait pod -l k8s-app=kube-dns -n kube-system --for condition=ready --timeout 30s

echo "$(date)  flux bootstrap" >> status

cat << "EOF" > flux-setup.sh
#!/bin/bash

sleep_seconds=30
status_code=1
retry_count=0

until [ $status_code == 0 ]; do

    if [ $retry_count -gt 0 ]
    then
      echo "$(date)  retrying flux bootstrap - $retry_count" >> status
      echo "retrying flux bootstrap - $retry_count"
      sleep $sleep_seconds
    fi

    retry_count=$((retry_count + 1))

    flux bootstrap git \
    --url https://github.com/retaildevcrews/red-gitops \
    --password $(cat /home/akdc/.ssh/akdc.pat) \
    --token-auth true \
    --path ./bootstrap

    status_code=$?

    echo "flux status code: $status_code"
done

retry_count=$((retry_count - 1))
echo "flux retries: $retry_count" >> /home/akdc/status

flux create source git gitops \
--url https://github.com/retaildevcrews/red-gitops \
--branch main \
--password $(cat /home/akdc/.ssh/akdc.pat) \

flux create kustomization bootstrap \
--source GitRepository/gitops \
--path ./bootstrap/east-nc-raleigh-102 \
--prune true \
--interval 1m

flux create kustomization apps \
--source GitRepository/gitops \
--path ./deploy/east-nc-raleigh-102 \
--prune true \
--interval 1m

flux reconcile source git gitops

kubectl get pods -A

EOF

# change ownership
chmod +x flux-setup.sh
chmod 600 ~/.ssh/akdc.pat
chown -R ${ME}:${ME} /home/${ME}

# create the tls secret
# this has to be installed before flux
kubectl create secret tls ssl-cert --cert .ssh/certs.pem --key .ssh/certs.key

# remove the certs
rm -f .ssh/certs.pem
rm -f .ssh/certs.key

# setup flux
#./flux-setup.sh

echo "$(date)  complete" >> status
echo "complete" >> status
