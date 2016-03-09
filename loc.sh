#!/usr/bin/env bash

# Counts line of code.
wc -l `find src/github.com/mikan -type f` | tail -1
