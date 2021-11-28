if (Get-Command 'go') {
    if (!(Test-Path '.\bin')) {
        New-Item -Path '.\bin' -ItemType Directory
    }
    if (!(Test-Path '.\bin\logs')) {
        New-Item -Path '.\bin\logs' -ItemType Directory
    }
    Invoke-Command {go build}
    Invoke-Command {.\soar.exe config setup --force}
} else {
    Write-Output 'Go v1.16 or above is required'
}
