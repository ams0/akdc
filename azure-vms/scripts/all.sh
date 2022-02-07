#!/bin/bash

cd ..

./create-cluster.sh central mo stlouis 101 centralus &
./create-cluster.sh central mo stlouis 102 centralus &
./create-cluster.sh central mo stlouis 103 centralus &

./create-cluster.sh central mo kc 101 centralus &
./create-cluster.sh central mo kc 102 centralus &
./create-cluster.sh central mo kc 103 centralus &

./create-cluster.sh central tx austin 101 centralus &
./create-cluster.sh central tx austin 102 centralus &
./create-cluster.sh central tx austin 103 centralus &

./create-cluster.sh central tx dallas 101 centralus &
./create-cluster.sh central tx dallas 102 centralus &
./create-cluster.sh central tx dallas 103 centralus &

./create-cluster.sh central tx houston 101 centralus &
./create-cluster.sh central tx houston 102 centralus &
./create-cluster.sh central tx houston 103 centralus &

./create-cluster.sh east ga atlanta 101 eastus &
./create-cluster.sh east ga atlanta 102 eastus &
./create-cluster.sh east ga atlanta 103 eastus &

./create-cluster.sh east ga athens 101 eastus &
./create-cluster.sh east ga athens 102 eastus &
./create-cluster.sh east ga athens 103 eastus &

./create-cluster.sh east nc charlotte 101 eastus &
./create-cluster.sh east nc charlotte 102 eastus &
./create-cluster.sh east nc charlotte 103 eastus &

./create-cluster.sh east nc raleigh 101 eastus &
./create-cluster.sh east nc raleigh 102 eastus &
./create-cluster.sh east nc raleigh 103 eastus &

./create-cluster.sh west ca la 101 westus &
./create-cluster.sh west ca la 102 westus &
./create-cluster.sh west ca la 103 westus &

./create-cluster.sh west ca sfo 101 westus &
./create-cluster.sh west ca sfo 102 westus &
./create-cluster.sh west ca sfo 103 westus &

./create-cluster.sh west ca sd 101 westus &
./create-cluster.sh west ca sd 102 westus &
./create-cluster.sh west ca sd 103 westus &

./create-cluster.sh west wa seattle 101 westus &
./create-cluster.sh west wa seattle 102 westus &
./create-cluster.sh west wa seattle 103 westus &
