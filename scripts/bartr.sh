#!/bin/bash

cd ..

./create-cluster.sh east ga athens 101 &
./create-cluster.sh east ga athens 102 &

./create-cluster.sh east ga atlanta 101 &
./create-cluster.sh east ga atlanta 102 &

./create-cluster.sh central tx austin 101 &
./create-cluster.sh central tx austin 102 &

./create-cluster.sh central tx dallas 101 &
./create-cluster.sh central tx dallas 102 &

./create-cluster.sh west wa seattle 101 &
./create-cluster.sh west wa seattle 102 &
