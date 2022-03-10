#!/bin/bash

# this runs each time the container starts

echo "$(date +'%Y-%m-%d %H:%M:%S')    post-start start" >> "$HOME/status"

# update the base docker images
docker pull mcr.microsoft.com/dotnet/sdk:5.0-alpine
docker pull mcr.microsoft.com/dotnet/aspnet:5.0-alpine
docker pull mcr.microsoft.com/dotnet/sdk:5.0
docker pull mcr.microsoft.com/dotnet/aspnet:6.0-alpine
docker pull mcr.microsoft.com/dotnet/sdk:6.0
docker pull ghcr.io/cse-labs/webv-red:latest
docker pull ghcr.io/cse-labs/webv-red:beta
docker pull ghcr.io/retaildevcrews/autogitops:beta

# build the cli
export KIC_PATH=/workspaces/akdc/bin
export KIC_NAME=akdc

cd src/kic || exit
make build
cd ../..

echo "$(date +'%Y-%m-%d %H:%M:%S')    post-start complete" >> "$HOME/status"
