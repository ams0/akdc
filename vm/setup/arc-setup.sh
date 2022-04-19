#!/usr/bin/env bash

set -e

echo "$(date +'%Y-%m-%d %H:%M:%S')  arc-setup start" >> "/home/${AKDC_ME}/status"

echo "$(date +'%Y-%m-%d %H:%M:%S')  arc-setup complete" >> "/home/${AKDC_ME}/status"
