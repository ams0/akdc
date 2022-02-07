#!/bin/bash

cd ..

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
