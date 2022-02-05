#!/bin/bash

cd ..

# remove the IP list
rm -f ips

./create-cluster.sh east ga athens 104 &
./create-cluster.sh east ga athens 105 &

./create-cluster.sh east ga atlanta 104 &
./create-cluster.sh east ga atlanta 105 &

./create-cluster.sh central tx austin 104 &
./create-cluster.sh central tx austin 105 &

./create-cluster.sh central tx dallas 104 &
./create-cluster.sh central tx dallas 105 &

./create-cluster.sh west wa seattle 104 &
./create-cluster.sh west wa seattle 105 &
