#!/bin/bash

akdc delete central-ks-kc-101
akdc delete central-ks-kc-102
akdc delete central-ks-kc-103
akdc delete central-ks-kc-104
akdc delete central-ks-kc-105
akdc delete central-tx-aus-101
akdc delete central-tx-aus-102
akdc delete central-tx-aus-103
akdc delete central-tx-aus-104
akdc delete central-tx-aus-105
akdc delete east-ga-atl-101
akdc delete east-ga-atl-102
akdc delete east-ga-atl-103
akdc delete east-ga-atl-104
akdc delete east-ga-atl-105
akdc delete west-ca-sand-101
akdc delete west-ca-sand-102
akdc delete west-ca-sand-103
akdc delete west-ca-sand-104
akdc delete west-ca-sand-105
akdc delete west-wa-sea-101
akdc delete west-wa-sea-102
akdc delete west-wa-sea-103
akdc delete west-wa-sea-104
akdc delete west-wa-sea-105

rm -f "$(dirname "${BASH_SOURCE[0]}")/ips"
