@echo off

go version >nul 2>&1 && (
    cd ..
    if not exist bin (
        mkdir bin
    )
    if not exist "bin/logs" (
        mkdir "bin/logs"
    )
    go build
    soar.exe config setup --force
) || (
    echo Go v1.16 or above is required
)
