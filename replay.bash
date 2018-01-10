#!/usr/bin/env bash

set -o nounset
set -o errexit
set -o pipefail
# set -o xtrace

HOSTNAME=$1

curl "${HOSTNAME}"

echo ''
