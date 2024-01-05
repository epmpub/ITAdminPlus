
#clear file

Remove-Item .\forWin.exe -ErrorAction SilentlyContinue
Remove-Item .\forArm  -ErrorAction SilentlyContinue
Remove-Item .\forLinux  -ErrorAction SilentlyContinue
Remove-Item .\forMac  -ErrorAction SilentlyContinue

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