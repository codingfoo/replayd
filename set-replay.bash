#!/usr/bin/env bash

set -o nounset
set -o errexit

HOSTNAME=$1
PAYLOAD=$2

curl -XPOST "$HOSTNAME" -d "$PAYLOAD"
