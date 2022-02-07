#!/bin/bash

cd ..

# add DO servers
echo "east-nc-raleigh-104	167.71.181.26" > ips
echo "east-nc-raleigh-105	165.227.74.132" >> ips

./create-cluster.sh east ga atlanta 104 &
./create-cluster.sh east ga atlanta 105 &

./create-cluster.sh central tx austin 104 &
./create-cluster.sh central tx austin 105 &

./create-cluster.sh central mo kc 104 &
./create-cluster.sh central mo kc 105 &

./create-cluster.sh west ca sd 104 &
./create-cluster.sh west ca sd 105 &

./create-cluster.sh west wa seattle 104 &
./create-cluster.sh west wa seattle 105 &
