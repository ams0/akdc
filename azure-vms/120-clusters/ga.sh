#!/bin/bash

cd ..

./create-cluster.sh east ga athens 101 eastus &
./create-cluster.sh east ga athens 102 eastus &
./create-cluster.sh east ga athens 103 eastus &
./create-cluster.sh east ga athens 104 eastus &
./create-cluster.sh east ga athens 105 eastus &

./create-cluster.sh east ga atlanta 101 eastus &
./create-cluster.sh east ga atlanta 102 eastus &
./create-cluster.sh east ga atlanta 103 eastus &
./create-cluster.sh east ga atlanta 104 eastus &
./create-cluster.sh east ga atlanta 105 eastus &

./create-cluster.sh east ga north 101 eastus &
./create-cluster.sh east ga north 102 eastus &
./create-cluster.sh east ga north 103 eastus &
./create-cluster.sh east ga north 104 eastus &
./create-cluster.sh east ga north 105 eastus &

./create-cluster.sh east ga south 101 eastus &
./create-cluster.sh east ga south 102 eastus &
./create-cluster.sh east ga south 103 eastus &
./create-cluster.sh east ga south 104 eastus &
./create-cluster.sh east ga south 105 eastus &
