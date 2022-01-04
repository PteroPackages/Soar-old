# Soar Package Installer
# TODO: turn off cmd/cmdlet output
# TODO: handle execution better
# TODO: add to PATH env

if (Get-Command 'go') {
    Write-Host 'soar: checking package...'

    if (!(Test-Path '..\main.go')) {
        Write-Host 'soar: installer must be executed in package directory.'
        Exit 1
    }

    Write-Host "soar: config directories will be setup at 'C:\soar'."
    $res = Read-Host -Prompt '    : do you want to continue? (y/n) '
    switch ($res) {
        'y' {}
        'Y' {}
        Default {
            Exit 0
        }
    }

    Write-Host 'soar: starting installation...'

    if (!(Test-Path 'C:\soar')) {
        New-Item -Path 'C:\soar' -ItemType Directory
    }

    if (!(Test-Path 'C:\soar\bin')) {
        New-Item -Path 'C:\soar\bin' -ItemType Directory
    }

    if (!(Test-Path 'C:\soar\bin\logs')) {
        New-Item -Path 'C:\soar\bin\logs' -ItemType Directory
    }

    Write-Host 'soar: building packages...'
    Set-Location '..'
    Invoke-Command {go build}

    Write-Host 'soar: attempting config setup...'
    Invoke-Command {.\soar.exe config setup}
    Write-Host 'soar: successfully installed Soar CLI!'
} else {
    Write-Host 'Go v1.16 or above is required'
    Exit 1
}
