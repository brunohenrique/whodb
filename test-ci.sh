#!/usr/bin/env bash

set -e

PACKAGES=$(find . -name '*.go' -print0 | xargs -0 -n1 dirname | sort --unique)
for d in $PACKAGES; do
    go test -v -race -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
