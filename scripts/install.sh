#!/usr/bin/env bash

# Soar Package Installer
# TODO: handle execution better
# TODO: set as global command (?)

go version >/dev/null 2>&1 && {
    echo "soar: checking package..."

    if [[ ! -f "../main.go" ]]; then
        echo "soar: installer must be executed in package directory."
        exit 1
    fi

    echo "soar: config directories will be setup at '/soar'."
    read -t 30 -n 1 -p "    : do you want to continue? (y/n) " RES
    if [[ -ne $RES "y" && -ne $RES "Y" ]]; then
        exit 0
    fi

    echo "soar: starting installation..."

    if [[ ! -d /soar ]]; then
        mkdir /soar
    fi

    if [[ ! -d /soar/bin ]]; then
        mkdir /soar/bin
    fi

    if [[ ! -d /soar/bin/logs ]]; then
        mkdir /soar/bin/logs
    fi

    echo "soar: building packages..."
    cd .. && go build

    echo "soar: attempting config setup..."
    ./soar config setup
    echo "soar: successfully installed Soar CLI!"
} || {
    echo >&2 "soar: Go v1.16 or above is required"
    exit 1
}
