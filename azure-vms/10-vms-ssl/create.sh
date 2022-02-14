#!/bin/bash

# change to this directory
cd "$(dirname "${BASH_SOURCE[0]}")" || exit

create-cluster -c central-ks-kc-104 -q -l centralus -z cseretail.com --ssl &
create-cluster -c central-ks-kc-105 -q -l centralus -z cseretail.com --ssl &

create-cluster -c central-tx-aus-104 -q -l centralus -z cseretail.com --ssl &
create-cluster -c central-tx-aus-105 -q -l centralus -z cseretail.com --ssl &

create-cluster -c east-ga-atl-104 -q -l eastus -z cseretail.com --ssl &
create-cluster -c east-ga-atl-105 -q -l eastus -z cseretail.com --ssl &

create-cluster -c west-ca-sand-104 -q -l westus -z cseretail.com --ssl &
create-cluster -c west-ca-sand-105 -q -l westus -z cseretail.com --ssl &

create-cluster -c west-wa-sea-104 -q -l westus -z cseretail.com --ssl &
create-cluster -c west-wa-sea-105 -q -l westus -z cseretail.com --ssl &
