#!/bin/bash

### runs as akdc user

# this runs after arc-setup.sh

set -e

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-private-repos start" >> /home/akdc/status

# add commands here

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-private-repos complete" >> /home/akdc/status
