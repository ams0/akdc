#!/bin/bash

cd ..

./create-cluster.sh west wa olympia 101 westus &
./create-cluster.sh west wa olympia 102 westus &
./create-cluster.sh west wa olympia 103 westus &

./create-cluster.sh west wa seattle 101 westus &
./create-cluster.sh west wa seattle 102 westus &
./create-cluster.sh west wa seattle 103 westus &

./create-cluster.sh west wa spokane 101 westus &
./create-cluster.sh west wa spokane 102 westus &
./create-cluster.sh west wa spokane 103 westus &
