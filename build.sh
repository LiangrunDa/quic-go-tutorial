#!/bin/bash

# Version
git rev-parse HEAD > VERSION

# download go1.18.9
wget -c https://go.dev/dl/go1.18.9.linux-amd64.tar.gz 
tar -C /usr/local -xzf go1.18.9.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# build
go mod tidy
source ./scripts/build_client.sh
source ./scripts/build_server.sh
source ./scripts/build_optimize.sh

# Add all necessary files to the zip
# Do NOT remove the three scripts
zip -r artifact.zip VERSION setup-env.sh run-client.sh run-server.sh client server optimize_client optimize_server