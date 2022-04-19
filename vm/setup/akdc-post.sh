#!/bin/bash

### runs as akdc user

# this runs after flux-setup.sh

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-post start" >> "/home/${AKDC_ME}/status"

docker pull ghcr.io/cse-labs/webv-red:latest
docker pull ghcr.io/cse-labs/webv-red:beta

docker pull ghcr.io/cse-labs/heartbeat
docker pull ghcr.io/cse-labs/imdb-app

docker run -d --restart always -p 30082:8080 ghcr.io/cse-labs/heartbeat
docker run -d --restart always -p 30080:8080 ghcr.io/cse-labs/imdb-app

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-post complete" >> "/home/${AKDC_ME}/status"
