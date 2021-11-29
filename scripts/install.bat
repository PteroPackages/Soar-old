:: Soar Package Installer
:: TODO: handle excution better
:: TODO: add to PATH env

@echo off

go version >nul 2>&1 && (
    echo soar: checking package...
    if not exist main.go (
        echo soar: installer must be executed in package directory.
        exit 1
    )

    echo soar: starting installation...

    if not exist bin (
        mkdir bin
    )

    if not exist "bin/logs" (
        mkdir "bin/logs"
    )

    echo soar: building packages...
    go build

    echo soar: attempting config setup...
    soar config setup --force
    echo soar: successfully installed Soar CLI!
) || (
    echo soar: Go v1.16 or above is required
    exit 1
)
