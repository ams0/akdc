#!/bin/bash

cd ..

az group delete -y --no-wait -g  west-wa-east-101 &
az group delete -y --no-wait -g  west-wa-east-102 &
az group delete -y --no-wait -g  west-wa-east-103 &
az group delete -y --no-wait -g  west-wa-east-104 &
az group delete -y --no-wait -g  west-wa-east-105 &

az group delete -y --no-wait -g  west-wa-olympia-101 &
az group delete -y --no-wait -g  west-wa-olympia-102 &
az group delete -y --no-wait -g  west-wa-olympia-103 &
az group delete -y --no-wait -g  west-wa-olympia-104 &
az group delete -y --no-wait -g  west-wa-olympia-105 &

az group delete -y --no-wait -g  west-wa-seattle-101 &
az group delete -y --no-wait -g  west-wa-seattle-102 &
az group delete -y --no-wait -g  west-wa-seattle-103 &
az group delete -y --no-wait -g  west-wa-seattle-104 &
az group delete -y --no-wait -g  west-wa-seattle-105 &

az group delete -y --no-wait -g  west-wa-spokane-101 &
az group delete -y --no-wait -g  west-wa-spokane-102 &
az group delete -y --no-wait -g  west-wa-spokane-103 &
az group delete -y --no-wait -g  west-wa-spokane-104 &
az group delete -y --no-wait -g  west-wa-spokane-105 &
