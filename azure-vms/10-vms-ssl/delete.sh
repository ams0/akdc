#!/bin/bash

akdc delete central-ks-kc-104
akdc delete central-ks-kc-105
akdc delete central-tx-aus-104
akdc delete central-tx-aus-105
akdc delete east-ga-atl-104
akdc delete east-ga-atl-105
akdc delete west-ca-sand-104
akdc delete west-ca-sand-105
akdc delete west-wa-sea-104
akdc delete west-wa-sea-105

rm -f "$(dirname "${BASH_SOURCE[0]}")/ips"
