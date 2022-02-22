#!/bin/bash

# this runs after flux-setup.sh
# this does not if akdc create --debug is used

# runs as akdc user
# env variables defined in .bashrc
    # AKDC_CLUSTER
    # AKDC_REPO
    # AKDC_FQDN
    # AKDC_DEBUG

# change to this directory
cd "$(dirname "${BASH_SOURCE[0]}")" || exit

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-post start" >> status

# add your post script here

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-post complete" >> status
