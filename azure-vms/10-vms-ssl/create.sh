#!/bin/bash

# change to this directory
cd "$(dirname "${BASH_SOURCE[0]}")" || exit

akdc create -c central-ks-kc-104 -q -l centralus -z cseretail.com --ssl &
akdc create -c central-ks-kc-105 -q -l centralus -z cseretail.com --ssl &

akdc create -c central-tx-aus-104 -q -l centralus -z cseretail.com --ssl &
akdc create -c central-tx-aus-105 -q -l centralus -z cseretail.com --ssl &

akdc create -c east-ga-atl-104 -q -l eastus -z cseretail.com --ssl &
akdc create -c east-ga-atl-105 -q -l eastus -z cseretail.com --ssl &

akdc create -c west-ca-sand-104 -q -l westus -z cseretail.com --ssl &
akdc create -c west-ca-sand-105 -q -l westus -z cseretail.com --ssl &

akdc create -c west-wa-sea-104 -q -l westus -z cseretail.com --ssl &
akdc create -c west-wa-sea-105 -q -l westus -z cseretail.com --ssl &
