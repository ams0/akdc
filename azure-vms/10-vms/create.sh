#!/bin/bash

# change to this directory
cd "$(dirname "${BASH_SOURCE[0]}")" || exit

akdc create -q -c central-ks-kc-101 &
akdc create -q -c central-ks-kc-102 &
akdc create -q -c central-tx-aus-101 &
akdc create -q -c central-tx-aus-102 &
akdc create -q -c east-ga-atl-101 &
akdc create -q -c east-ga-atl-102 &
akdc create -q -c west-ca-sand-101 &
akdc create -q -c west-ca-sand-102 &
akdc create -q -c west-wa-sea-101 &
akdc create -q -c west-wa-sea-102 &
