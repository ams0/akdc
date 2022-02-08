#!/bin/bash

cd ..

./create-cluster.sh west wa east 101 westus &
./create-cluster.sh west wa east 102 westus &
./create-cluster.sh west wa east 103 westus &
./create-cluster.sh west wa east 104 westus &
./create-cluster.sh west wa east 105 westus &

./create-cluster.sh west wa olympia 101 westus &
./create-cluster.sh west wa olympia 102 westus &
./create-cluster.sh west wa olympia 103 westus &
./create-cluster.sh west wa olympia 104 westus &
./create-cluster.sh west wa olympia 105 westus &

./create-cluster.sh west wa seattle 101 westus &
./create-cluster.sh west wa seattle 102 westus &
./create-cluster.sh west wa seattle 103 westus &
./create-cluster.sh west wa seattle 104 westus &
./create-cluster.sh west wa seattle 105 westus &

./create-cluster.sh west wa spokane 101 westus &
./create-cluster.sh west wa spokane 102 westus &
./create-cluster.sh west wa spokane 103 westus &
./create-cluster.sh west wa spokane 104 westus &
./create-cluster.sh west wa spokane 105 westus &
