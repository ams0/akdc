#!/bin/bash

cd ..

./create-cluster.sh east nc charlotte 101 eastus &
./create-cluster.sh east nc charlotte 102 eastus &
./create-cluster.sh east nc charlotte 103 eastus &

./create-cluster.sh east nc durham 101 eastus &
./create-cluster.sh east nc durham 102 eastus &
./create-cluster.sh east nc durham 103 eastus &

./create-cluster.sh east nc west 101 eastus &
./create-cluster.sh east nc west 102 eastus &
./create-cluster.sh east nc west 103 eastus &
