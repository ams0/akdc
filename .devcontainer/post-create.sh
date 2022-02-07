#!/bin/bash

# this runs at Codespace creation - not part of pre-build

echo "$(date)    post-create start" >> ~/status

# clone repos
git -C /workspaces/ngsa pull
git -C /workspaces/webvalidate pull
git -C /workspaces/ngsa-app pull

echo "$(date)    post-create complete" >> ~/status
