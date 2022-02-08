#!/bin/bash

cd ..

./create-cluster.sh west ca la 101 westus &
./create-cluster.sh west ca la 102 westus &
./create-cluster.sh west ca la 103 westus &

./create-cluster.sh west ca sfo 101 westus &
./create-cluster.sh west ca sfo 102 westus &
./create-cluster.sh west ca sfo 103 westus &

./create-cluster.sh west ca sd 101 westus &
./create-cluster.sh west ca sd 102 westus &
./create-cluster.sh west ca sd 103 westus &
