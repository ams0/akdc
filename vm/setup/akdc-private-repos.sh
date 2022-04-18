#!/bin/bash

### runs as akdc user

# this runs after arc-setup.sh

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-private-repos start" >> "/home/${AKDC_ME}/status"

# add commands here

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-private-repos complete" >> "/home/${AKDC_ME}/status"
