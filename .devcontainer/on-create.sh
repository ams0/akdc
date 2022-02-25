#!/bin/bash

# this runs as part of pre-build

echo "$(date +'%Y-%m-%d %H:%M:%S')    on-create start" >> "$HOME/status"

# do this early to avoid the popup
dotnet restore src/gen-gitops

export REPO_BASE=$PWD

mkdir -p "$HOME/.ssh"
mkdir -p "$HOME/.oh-my-zsh/completions"

# add cli completions
cp src/_* "$HOME/.oh-my-zsh/completions"

{
    # add cli to path
    echo "export PATH=\$PATH:$REPO_BASE/bin"

    # create aliases
    # make akdc
    echo "alias ma='cd $REPO_BASE/src/akdc && make; cd \$OLDPWD'"

    # make kic
    echo "alias mk='cd $REPO_BASE/src/kic && make; cd \$OLDPWD'"

    echo "export REPO_BASE=$PWD"
    echo "compinit"

    echo "export GOPATH=\$HOME/go"
} >> "$HOME/.zshrc"

# copy grafana.db to /grafana
sudo cp inner-loop/grafanadata/grafana.db /grafana
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

# install cobra
go install github.com/spf13/cobra/cobra@latest

# install golint
go install golang.org/x/lint/golint@latest

# clone repos
pushd ..
git clone https://github.com/microsoft/webvalidate
git clone https://github.com/retaildevcrews/ngsa-app
git clone https://github.com/retaildevcrews/edge-apps
git clone https://github.com/retaildevcrews/edge-gitops
git clone https://github.com/retaildevcrews/red-apps
git clone https://github.com/retaildevcrews/red-gitops
popd || exit

echo "$(date +'%Y-%m-%d %H:%M:%S')    on-create complete" >> "$HOME/status"
