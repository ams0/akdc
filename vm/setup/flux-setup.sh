#!/bin/bash

### runs as akdc user

echo "$(date +'%Y-%m-%d %H:%M:%S')  flux bootstrap start" >> "/home/${AKDC_ME}/status"

echo "$(date +'%Y-%m-%d %H:%M:%S')  flux bootstrap complete" >> "/home/${AKDC_ME}/status"
