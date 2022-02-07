#!/bin/bash

cd ..

./create-cluster.sh west wa seattle 101 westus &
./create-cluster.sh west wa seattle 102 westus &
./create-cluster.sh west wa seattle 103 westus &
