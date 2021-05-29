#!/usr/bin/env bash

# Show env vars
grep -v '^#' hide_seek.env

# Export env vars
export $(grep -v '^#' hide_seek.env | xargs)