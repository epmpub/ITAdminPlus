
#clear file

Remove-Item .\forWin.exe -ErrorAction SilentlyContinue
Remove-Item .\forArm  -ErrorAction SilentlyContinue
Remove-Item .\forLinux  -ErrorAction SilentlyContinue
Remove-Item .\forMac  -ErrorAction SilentlyContinue
Remove-Item .\forMac  -ErrorAction SilentlyContinue
Remove-Item .\md5.txt  -ErrorAction SilentlyContinue


$env:CGO_ENABLED="0"
$env:GOOS="linux"
$env:GOARCH="amd64"


go build -o forLinux


$env:CGO_ENABLED="0"
$env:GOOS="linux"
$env:GOARCH="arm64"

go build -o forArm

$env:CGO_ENABLED="0"
$env:GOOS="windows"
$env:GOARCH="amd64"
go build -o forWin.exe


$env:CGO_ENABLED="0"
$env:GOOS="darwin"
$env:GOARCH="amd64"
go build -o forMac

$fileInfo = [ordered]@{}

$bin = Get-FileHash -Algorithm MD5 .\forWin.exe
$timestamp = Get-Date -f 'o'

$fileInfo["Hash"] = $bin.Hash
$fileInfo["TimeStamp"] = $timestamp
$fileInfo["Version"] = Get-Random -Maximum 100

$fileInfo | ConvertTo-Json | Out-File .\forWin_md5.json

ossutil.exe cp .\forWin.exe oss://it2u/forWin.exe -f
ossutil.exe cp .\forWin_md5.json oss://it2u/forWin_md5.json -f

ossutil.exe set-acl oss://it2u/forWin.exe public-read
ossutil.exe set-acl oss://it2u/forWin_md5.json public-read
