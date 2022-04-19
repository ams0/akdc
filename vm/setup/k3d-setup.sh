#!/bin/bash

### if running manually, you must start a new shell to get the docker permissions

# change to this directory
cd "$(dirname "${BASH_SOURCE[0]}")" || exit

echo "$(date +'%Y-%m-%d %H:%M:%S')  k3d-setup start" >> "/home/${AKDC_ME}/status"

echo "$(date +'%Y-%m-%d %H:%M:%S')  k3d-setup complete" >> "/home/${AKDC_ME}/status"
