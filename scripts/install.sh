#!/usr/bin/env bash

go version >/dev/null 2>&1 && {
    cd ..
    if [[ ! -d bin ]]; then
        mkdir bin
    fi
    if [[ ! -d /bin/logs ]]; then
        mkdir /bin/logs
    fi
    go build
    ./soar config setup --force
} || {
    echo >&2 "Go v1.16 or above is required"
}
