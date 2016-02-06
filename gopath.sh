#!/usr/bin/env bash

# Update GOPATH by current directory.

export GOPATH="$(cd "$(dirname "${BASH_SOURCE:-$0}")"; pwd)"
echo GOPATH updated to $GOPATH
export PATH=$PATH:$GOPATH/bin
