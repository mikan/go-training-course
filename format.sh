#!/usr/bin/env bash

# Formats all go sources.

find src/github.com/mikan -name "*.go" -exec bin/goimports -l -w {} \;
