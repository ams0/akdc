#!/bin/bash

# change to this directory
cd "$(dirname "${BASH_SOURCE[0]}")" || exit

akdc create -c central-ks-kc-101 -q -l centralus &
akdc create -c central-ks-kc-102 -q -l centralus &
akdc create -c central-tx-aus-101 -q -l centralus &
akdc create -c central-tx-aus-102 -q -l centralus &
akdc create -c east-ga-atl-101 -q -l eastus &
akdc create -c east-ga-atl-102 -q -l eastus &
akdc create -c west-ca-sand-101 -q -l westus &
akdc create -c west-ca-sand-102 -q -l westus &
akdc create -c west-wa-sea-101 -q -l westus &
akdc create -c west-wa-sea-102 -q -l westus &
