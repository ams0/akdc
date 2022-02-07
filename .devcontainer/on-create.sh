#!/bin/bash

# this runs as part of pre-build

echo "$(date)    on-create start" >> ~/status

# do this early to avoid the popup
dotnet restore src/gen-gitops

# add cli to path
echo 'export PATH=$PATH:/workspaces/akdc/src/cli' >> $HOME/.zshrc

# add akdc completions
cp src/cli/_akdc $HOME/.oh-my-zsh/completions
unfunction _akdc && autoload -Uz _akdc

# clone repos
git clone https://github.com/retaildevcrews/edge-ngsa /workspaces/ngsa
git clone https://github.com/microsoft/webvalidate /workspaces/webvalidate
git clone https://github.com/retaildevcrews/ngsa-app /workspaces/ngsa-app

# copy grafana.db to /grafana
sudo cp inner-loop/grafanadata/grafana.db /grafana
sudo chown -R 472:0 /grafana

# create local registry
docker network create k3d
k3d registry create registry.localhost --port 5500
docker network connect k3d k3d-registry.localhost

# update the base docker images
docker pull mcr.microsoft.com/dotnet/sdk:5.0-alpine
docker pull mcr.microsoft.com/dotnet/aspnet:5.0-alpine
docker pull mcr.microsoft.com/dotnet/sdk:5.0
docker pull mcr.microsoft.com/dotnet/aspnet:6.0-alpine
docker pull mcr.microsoft.com/dotnet/sdk:6.0

echo "$(date)    on-create complete" >> ~/status
