#!/usr/bin/env bash

# Update GOPATH by current directory.
# Usage: source gopath.sh

export GOPATH=$GOPATH:"$(cd "$(dirname "${BASH_SOURCE:-$0}")"; pwd)"
echo GOPATH updated to $GOPATH
export PATH=$PATH:$GOPATH/bin
