#!/usr/bin/env bash

set -o nounset
set -o errexit
set -o pipefail
# set -o xtrace

go run -race main.go &

wrk -t12 -c400 -d30s http://127.0.0.1:8080/ &
wrk -t12 -c400 -d30s -s post.lua http://127.0.0.1:8080/ &
