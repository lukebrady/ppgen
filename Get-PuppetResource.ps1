function Get-PuppetGeneratorHealth {
    [CmdletBinding()]
    param (
        [String] $URI
    )

    $(Invoke-WebRequest -Method Get -Uri $URI).Content
}

function Get-PuppetResource {
    [CmdletBinding()]
    param 
    (
        [String]    $URI,
        [String[]]  $Resources,
        [String]    $PuppetFile
    )

    if(Test-Path $PuppetFile) {
        Write-Output "Puppet File already exists.`n"
    }

    else {
        New-Item -Path $PuppetFile -ItemType File
    }

    try {
        foreach($Resource in $Resources) {
            $(Invoke-WebRequest -Method Get -Uri "$URI/$($Resource.toLower())").Content >> $PuppetFile
        }
    } catch { $Error }

    if($LASTEXITCODE -eq 0) {
        Write-Host "$PuppetFile was created successfully.`n"
    }
    else {
        "There was an error creating $PuppetFile.`n"
    }

}

# Get-PuppetGeneratorHealth -URI "http://localhost:8080"
Get-PuppetResource -URI "http://67.205.135.251" -Resources "file","user","group" -PuppetFile "C:\Users\lbrad23105\Dropbox\Software\Go\ppgen\wow.pp"