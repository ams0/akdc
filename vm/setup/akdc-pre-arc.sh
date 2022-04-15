#!/bin/bash

### runs as akdc user

# this runs before arc-setup.sh

set -e

# change to this directory
# cd "$(dirname "${BASH_SOURCE[0]}")" || exit

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre-arc start" >> /home/akdc/status

# add commands here

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre-arc complete" >> /home/akdc/status
