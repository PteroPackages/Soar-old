# Soar Package Installer
# TODO: turn off cmd/cmdlet output
# TODO: handle execution better
# TODO: add to PATH env

if (Get-Command 'go') {
    Write-Output 'soar: checking package...'
    if (!(Test-Path '.\main.go')) {
        Write-Output 'soar: installer must be executed in package directory'
        Exit 1
    }

    Write-Output 'soar: starting installation...'

    if (!(Test-Path '.\bin')) {
        New-Item -Path '.\bin' -ItemType Directory
    }

    if (!(Test-Path '.\bin\logs')) {
        New-Item -Path '.\bin\logs' -ItemType Directory
    }

    Write-Output 'soar: building packages...'
    Invoke-Command {go build}

    Write-Output 'soar: attempting config setup...'
    Invoke-Command {.\soar.exe config setup --force}
    Write-Output 'soar: successfully installed Soar CLI!'
} else {
    Write-Output 'Go v1.16 or above is required'
    Exit 1
}
