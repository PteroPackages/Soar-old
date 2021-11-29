#!/usr/bin/env bash

# Soar Package Installer
# TODO: handle execution better
# TODO: set as global command (?)

go version >/dev/null 2>&1 && {
    echo "soar: checking package..."
    if [[ ! -f main.go ]]; then
        echo "soar: installer must be executed in package directory."
        exit 1
    fi

    echo "soar: starting installation..."

    if [[ ! -d bin ]]; then
        mkdir bin
    fi

    if [[ ! -d /bin/logs ]]; then
        mkdir /bin/logs
    fi

    echo "soar: building packages..."
    go build

    echo "soar: attempting config setup..."
    ./soar config setup --force
    echo "soar: successfully installed Soar CLI!"
} || {
    echo >&2 "soar: Go v1.16 or above is required"
    exit 1
}
