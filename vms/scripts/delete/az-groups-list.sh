#!/bin/bash

az group list -o table | sort | grep -e central- -e east- -e west-
