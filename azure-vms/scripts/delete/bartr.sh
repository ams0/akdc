#!/bin/bash

az group delete -y --no-wait -g central-tx-austin-104
az group delete -y --no-wait -g central-tx-austin-105
az group delete -y --no-wait -g central-tx-dallas-104
az group delete -y --no-wait -g central-tx-dallas-105

az group delete -y --no-wait -g east-ga-atlanta-104
az group delete -y --no-wait -g east-ga-atlanta-105
az group delete -y --no-wait -g east-ga-athens-104
az group delete -y --no-wait -g east-ga-athens-105

az group delete -y --no-wait -g west-wa-seattle-104
az group delete -y --no-wait -g west-wa-seattle-105

az group list -o table | sort | grep -e -10
