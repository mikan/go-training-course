#!/usr/bin/env bash

# Formats all go sources.

find ch* -name "*.go" -exec ../../../../bin/goimports -l -w {} \;
