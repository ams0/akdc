#!/bin/bash

cd ..

az group delete -y --no-wait -g  central-tx-austin-101 &
az group delete -y --no-wait -g  central-tx-austin-102 &
az group delete -y --no-wait -g  central-tx-austin-103 &
az group delete -y --no-wait -g  central-tx-austin-104 &
az group delete -y --no-wait -g  central-tx-austin-105 &

az group delete -y --no-wait -g  central-tx-dallas-101 &
az group delete -y --no-wait -g  central-tx-dallas-102 &
az group delete -y --no-wait -g  central-tx-dallas-103 &
az group delete -y --no-wait -g  central-tx-dallas-104 &
az group delete -y --no-wait -g  central-tx-dallas-105 &

az group delete -y --no-wait -g  central-tx-houston-101 &
az group delete -y --no-wait -g  central-tx-houston-102 &
az group delete -y --no-wait -g  central-tx-houston-103 &
az group delete -y --no-wait -g  central-tx-houston-104 &
az group delete -y --no-wait -g  central-tx-houston-105 &

az group delete -y --no-wait -g  central-tx-san-101 &
az group delete -y --no-wait -g  central-tx-san-102 &
az group delete -y --no-wait -g  central-tx-san-103 &
az group delete -y --no-wait -g  central-tx-san-104 &
az group delete -y --no-wait -g  central-tx-san-105 &
