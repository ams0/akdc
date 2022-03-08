#!/bin/bash

# this runs as part of pre-build

echo "$(date +'%Y-%m-%d %H:%M:%S')    on-create start" >> "$HOME/status"

export REPO_BASE=$PWD
export AKDC_REPO=retaildevcrews/edge-gitops

export PATH="$PATH:$REPO_BASE/bin"
export GOPATH="$HOME/go"

mkdir -p .ssh

mkdir -p "$HOME/.ssh"
mkdir -p "$HOME/.oh-my-zsh/completions"

{
    echo "defaultIPs: /workspaces/akdc/red-fleet/ips"
    echo "reservedClusterPrefixes: corp-monitoring central-mo-kc central-tx-austin east-ga-atlanta east-nc-raleigh west-ca-sd west-wa-redmond west-wa-seattle"
} > "$HOME/.kic"

# add cli completions
cp src/_* "$HOME/.oh-my-zsh/completions"

{
    #shellcheck disable=2016,2028
    echo 'hsort() { read -r; printf "%s\\n" "$REPLY"; sort }'

    # add cli to path
    echo "export PATH=\$PATH:$REPO_BASE/bin"
    echo "export GOPATH=\$HOME/go"

    # create aliases
    echo "alias mk='cd $REPO_BASE/src/kic && make build; cd \$OLDPWD'"
    echo "alias akdc='kic manage fleet'"
    echo "alias flt='kic manage fleet'"

    echo "export REPO_BASE=$PWD"
    echo "export AKDC_REPO=retaildevcrews/edge-gitops"
    echo "export AKDC_SSL=cseretail.com"
    echo "export AKDC_GITOPS=true"
    echo "export KIC_PATH=/workspaces/akdc/bin"
    echo "export KIC_NAME=kic"
    echo "compinit"
} >> "$HOME/.zshrc"

# copy grafana.db to /grafana
sudo cp .devcontainer/grafana.db /grafana
sudo chown -R 472:0 /grafana

# create local registry
docker network create k3d
k3d registry create registry.localhost --port 5500
docker network connect k3d k3d-registry.localhost

# pull the base docker images
docker pull mcr.microsoft.com/dotnet/sdk:5.0-alpine
docker pull mcr.microsoft.com/dotnet/aspnet:5.0-alpine
docker pull mcr.microsoft.com/dotnet/sdk:5.0
docker pull mcr.microsoft.com/dotnet/aspnet:6.0-alpine
docker pull mcr.microsoft.com/dotnet/sdk:6.0
docker pull ghcr.io/cse-labs/webv-red:latest
docker pull ghcr.io/cse-labs/webv-red:beta
docker pull ghcr.io/retaildevcrews/autogitops:beta

# install cobra
go install -v github.com/spf13/cobra/cobra@latest

# install go modules
go install -v golang.org/x/lint/golint@latest
go install -v github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest
go install -v github.com/ramya-rao-a/go-outline@latest
go install -v github.com/cweill/gotests/gotests@latest
go install -v github.com/fatih/gomodifytags@latest
go install -v github.com/josharian/impl@latest
go install -v github.com/haya14busa/goplay/cmd/goplay@latest
go install -v github.com/go-delve/delve/cmd/dlv@latest
go install -v honnef.co/go/tools/cmd/staticcheck@latest
go install -v golang.org/x/tools/gopls@latest

cp -r .devcontainer/.vscode .

# clone repos
cd ..
git clone https://github.com/microsoft/webvalidate
git clone https://github.com/retaildevcrews/ngsa-app
git clone https://github.com/retaildevcrews/edge-apps
git clone https://github.com/retaildevcrews/edge-gitops
git clone https://github.com/retaildevcrews/red-apps
git clone https://github.com/retaildevcrews/red-gitops
cd "$OLDPWD" || exit

echo "$(date +'%Y-%m-%d %H:%M:%S')    on-create complete" >> "$HOME/status"
