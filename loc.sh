#!/usr/bin/env bash

# Counts line of code.
wc -l `find ch* -type f` | tail -1
