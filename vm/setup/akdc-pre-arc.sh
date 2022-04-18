#!/bin/bash

### runs as akdc user

# this runs before arc-setup.sh

# change to this directory
# cd "$(dirname "${BASH_SOURCE[0]}")" || exit

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre-arc start" >> "/home/${AKDC_ME}/status"

# add commands here

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre-arc complete" >> "/home/${AKDC_ME}/status"
