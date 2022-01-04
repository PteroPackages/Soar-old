:: Soar Package Installer
:: TODO: handle excution better
:: TODO: add to PATH env

@echo off

go version >nul 2>&1 (
    echo soar: checking package...

    if not exist "..\main.go" (
        echo soar: installer must be executed in package directory.
        exit 1
    )

    echo soar: config directories will be setup at 'C:\soar'.
    choice /T 30 /D N /N /M "    : do you want to continue? (y/n) "

    if %ERRORLEVEL% neq 0 (
        exit 0
    )

    echo soar: starting installation...

    if not exist "C:\soar" (
        mkdir "C:\soar"
    )

    if not exist "C:\soar\bin" (
        mkdir "C:\soar\bin"
    )

    if not exist "C:\soar\bin\logs" (
        mkdir "C:\soar\bin\logs"
    )

    echo soar: building packages...
    cd .. && go build

    echo soar: attempting config setup...
    soar config setup
    echo soar: successfully installed Soar CLI!
) || (
    echo soar: Go v1.16 or above is required
    exit 1
)
