#!/bin/bash

cd ..

az group delete -y --no-wait -g  central-mo-columbia-101 &
az group delete -y --no-wait -g  central-mo-columbia-102 &
az group delete -y --no-wait -g  central-mo-columbia-103 &
az group delete -y --no-wait -g  central-mo-columbia-104 &
az group delete -y --no-wait -g  central-mo-columbia-105 &

az group delete -y --no-wait -g  central-mo-east-101 &
az group delete -y --no-wait -g  central-mo-east-102 &
az group delete -y --no-wait -g  central-mo-east-103 &
az group delete -y --no-wait -g  central-mo-east-104 &
az group delete -y --no-wait -g  central-mo-east-105 &

az group delete -y --no-wait -g  central-mo-kc-101 &
az group delete -y --no-wait -g  central-mo-kc-102 &
az group delete -y --no-wait -g  central-mo-kc-103 &
az group delete -y --no-wait -g  central-mo-kc-104 &
az group delete -y --no-wait -g  central-mo-kc-105 &

az group delete -y --no-wait -g  central-mo-north-101 &
az group delete -y --no-wait -g  central-mo-north-102 &
az group delete -y --no-wait -g  central-mo-north-103 &
az group delete -y --no-wait -g  central-mo-north-104 &
az group delete -y --no-wait -g  central-mo-north-105 &

az group delete -y --no-wait -g  central-mo-south-101 &
az group delete -y --no-wait -g  central-mo-south-102 &
az group delete -y --no-wait -g  central-mo-south-103 &
az group delete -y --no-wait -g  central-mo-south-104 &
az group delete -y --no-wait -g  central-mo-south-105 &

az group delete -y --no-wait -g  central-mo-stlouis-101 &
az group delete -y --no-wait -g  central-mo-stlouis-102 &
az group delete -y --no-wait -g  central-mo-stlouis-103 &
az group delete -y --no-wait -g  central-mo-stlouis-104 &
az group delete -y --no-wait -g  central-mo-stlouis-105 &

az group delete -y --no-wait -g  central-mo-west-101 &
az group delete -y --no-wait -g  central-mo-west-102 &
az group delete -y --no-wait -g  central-mo-west-103 &
az group delete -y --no-wait -g  central-mo-west-104 &
az group delete -y --no-wait -g  central-mo-west-105 &