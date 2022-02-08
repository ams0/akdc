#!/bin/bash

cd ..

az group delete -y --no-wait -g  west-ca-la-101 &
az group delete -y --no-wait -g  west-ca-la-102 &
az group delete -y --no-wait -g  west-ca-la-103 &
az group delete -y --no-wait -g  west-ca-la-104 &
az group delete -y --no-wait -g  west-ca-la-105 &

az group delete -y --no-wait -g  west-ca-sac-101 &
az group delete -y --no-wait -g  west-ca-sac-102 &
az group delete -y --no-wait -g  west-ca-sac-103 &
az group delete -y --no-wait -g  west-ca-sac-104 &
az group delete -y --no-wait -g  west-ca-sac-105 &

az group delete -y --no-wait -g  west-ca-sfo-101 &
az group delete -y --no-wait -g  west-ca-sfo-102 &
az group delete -y --no-wait -g  west-ca-sfo-103 &
az group delete -y --no-wait -g  west-ca-sfo-104 &
az group delete -y --no-wait -g  west-ca-sfo-105 &

az group delete -y --no-wait -g  west-ca-sd-101 &
az group delete -y --no-wait -g  west-ca-sd-102 &
az group delete -y --no-wait -g  west-ca-sd-103 &
az group delete -y --no-wait -g  west-ca-sd-104 &
az group delete -y --no-wait -g  west-ca-sd-105 &
