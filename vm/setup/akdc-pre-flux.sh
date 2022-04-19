#!/bin/bash

# this runs before flux-setup.sh

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre-flux start" >> "/home/${AKDC_ME}/status"

# change to this directory
#cd "$(dirname "${BASH_SOURCE[0]}")" || exit

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre-flux complete" >> "/home/${AKDC_ME}/status"
