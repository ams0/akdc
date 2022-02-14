#!/bin/bash

# change to this directory
cd "$(dirname "${BASH_SOURCE[0]}")" || exit

create-cluster -c central-ks-kc-101 -q -l centralus &
create-cluster -c central-ks-kc-102 -q -l centralus &
create-cluster -c central-tx-aus-101 -q -l centralus &
create-cluster -c central-tx-aus-102 -q -l centralus &
create-cluster -c east-ga-atl-101 -q -l eastus &
create-cluster -c east-ga-atl-102 -q -l eastus &
create-cluster -c west-ca-sand-101 -q -l westus &
create-cluster -c west-ca-sand-102 -q -l westus &
create-cluster -c west-wa-sea-101 -q -l westus &
create-cluster -c west-wa-sea-102 -q -l westus &
