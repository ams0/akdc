#!/bin/bash

echo "on-create start" >> ~/status

# do this early to avoid the popup
dotnet restore src/gen-gitops

echo "on-create complete" >> ~/status
