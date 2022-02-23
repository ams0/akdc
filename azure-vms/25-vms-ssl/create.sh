#!/bin/bash

# change to this directory
cd "$(dirname "${BASH_SOURCE[0]}")" || exit

akdc create -q --ssl cseretail.com -c central-ks-kc-101 &
akdc create -q --ssl cseretail.com -c central-ks-kc-102 &
akdc create -q --ssl cseretail.com -c central-ks-kc-103 &
akdc create -q --ssl cseretail.com -c central-ks-kc-104 &
akdc create -q --ssl cseretail.com -c central-ks-kc-105 &

akdc create -q --ssl cseretail.com -c central-tx-aus-101 &
akdc create -q --ssl cseretail.com -c central-tx-aus-102 &
akdc create -q --ssl cseretail.com -c central-tx-aus-103 &
akdc create -q --ssl cseretail.com -c central-tx-aus-104 &
akdc create -q --ssl cseretail.com -c central-tx-aus-105 &

akdc create -q --ssl cseretail.com -c east-ga-atl-101 &
akdc create -q --ssl cseretail.com -c east-ga-atl-102 &
akdc create -q --ssl cseretail.com -c east-ga-atl-103 &
akdc create -q --ssl cseretail.com -c east-ga-atl-104 &
akdc create -q --ssl cseretail.com -c east-ga-atl-105 &

akdc create -q --ssl cseretail.com -c west-ca-sand-101 &
akdc create -q --ssl cseretail.com -c west-ca-sand-102 &
akdc create -q --ssl cseretail.com -c west-ca-sand-103 &
akdc create -q --ssl cseretail.com -c west-ca-sand-104 &
akdc create -q --ssl cseretail.com -c west-ca-sand-105 &

akdc create -q --ssl cseretail.com -c west-wa-sea-101 &
akdc create -q --ssl cseretail.com -c west-wa-sea-102 &
akdc create -q --ssl cseretail.com -c west-wa-sea-103 &
akdc create -q --ssl cseretail.com -c west-wa-sea-104 &
akdc create -q --ssl cseretail.com -c west-wa-sea-105 &
