#!/bin/bash

# this runs as part of pre-build

echo "$(date)    on-create start" >> "$HOME/status"

# do this early to avoid the popup
dotnet restore src/gen-gitops

{
    # add cli to path
    echo "export PATH=\$PATH:/workspaces/akdc/src/cli"
    echo "alias mk='cd /workspaces/akdc/src/go-cli && make && cd \$OLDPWD'"

    # todo - hot fix - this is set to /go upstream
    echo "export GOPATH=\$HOME/go"
} >> "$HOME/.zshrc"

# add akdc completions
cp src/cli/_akdc "$HOME/.oh-my-zsh/completions"
unfunction _akdc
autoload -Uz _akdc
compinit

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

# install cobra
go install github.com/spf13/cobra/cobra@latest

# install golint
go install golang.org/x/lint/golint@latest

echo "$(date)    on-create complete" >> "$HOME/status"
