#!/usr/bin/env bash

# Update GOPATH by current directory.
# Usage: source gopath.sh

current="$(cd "$(dirname "${BASH_SOURCE:-$0}")"; pwd)"

if [ -z "$GOPATH" ]; then
    export GOPATH=$current
else
    export GOPATH=$GOPATH:$current
fi
echo GOPATH updated to $GOPATH
export PATH=$PATH:$GOPATH/bin
