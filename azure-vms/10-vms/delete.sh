#!/bin/bash

az group delete -y --no-wait -g central-tx-aus-101 &
az group delete -y --no-wait -g central-tx-aus-102 &
az group delete -y --no-wait -g central-ks-kc-101 &
az group delete -y --no-wait -g central-ks-kc-102 &
az group delete -y --no-wait -g east-ga-atl-101 &
az group delete -y --no-wait -g east-ga-atl-102 &
az group delete -y --no-wait -g west-ca-sand-101 &
az group delete -y --no-wait -g west-ca-sand-102 &
az group delete -y --no-wait -g west-wa-sea-101 &
az group delete -y --no-wait -g west-wa-sea-102 &

rm -f "$(dirname "${BASH_SOURCE[0]}")/ips"
