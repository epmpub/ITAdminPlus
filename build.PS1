
#clear file

del .\forWin.exe
del .\forArm
del .\forLinux

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