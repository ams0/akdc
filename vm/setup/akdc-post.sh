#!/bin/bash

### runs as akdc user

# this runs after flux-setup.sh

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-post start" >> "/home/${AKDC_ME}/status"

docker pull ghcr.io/cse-labs/webv-red:latest
docker pull ghcr.io/cse-labs/webv-red:beta

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-post complete" >> "/home/${AKDC_ME}/status"
