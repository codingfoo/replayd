#!/usr/bin/env bash

set -o nounset
set -o errexit

HOSTNAME=$1

curl "$HOSTNAME"

echo ''
