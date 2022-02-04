#!/bin/bash

az group delete -y --no-wait -g central-tx-austin-101
az group delete -y --no-wait -g central-tx-austin-102
az group delete -y --no-wait -g central-tx-dallas-101
az group delete -y --no-wait -g central-tx-dallas-102

az group delete -y --no-wait -g east-ga-atlanta-101
az group delete -y --no-wait -g east-ga-atlanta-102
az group delete -y --no-wait -g east-ga-athens-101
az group delete -y --no-wait -g east-ga-athens-102

az group delete -y --no-wait -g west-wa-seattle-101
az group delete -y --no-wait -g west-wa-seattle-102

az group list -o table | sort | grep -e central- -e east- -e west-
