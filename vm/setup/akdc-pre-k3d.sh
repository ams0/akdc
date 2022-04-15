#!/bin/bash

### runs as akdc user

# this runs before k3d-setup.sh

set -e

# change to this directory
#cd "$(dirname "${BASH_SOURCE[0]}")" || exit

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre-k3d start" >> /home/akdc/status

# add commands here

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre-k3d complete" >> /home/akdc/status
