# Arm Board
Write-Host -BackgroundColor Yellow -ForegroundColor White "Copy to ARM"
ssh openailab@192.168.3.146 "killall forArm"
scp forArm openailab@192.168.3.146:~
ssh openailab@192.168.3.146 "nohup ./forArm > /dev/null 2>&1 &"


# utools
Write-Host -BackgroundColor Yellow -ForegroundColor White "Copy to utools"
ssh root@39.108.212.138 "killall forLinux"
scp forLinux root@39.108.212.138:~
ssh root@39.108.212.138 "nohup ./forLinux > /dev/null 2>&1 &"

# WebServ
Write-Host -BackgroundColor Yellow -ForegroundColor White "Copy to WebServ"
ssh root@120.79.203.54 "killall forLinux"
scp forLinux root@120.79.203.54:~
ssh root@120.79.203.54 "nohup ./forLinux > /dev/null 2>&1 &"

# PrivateServ
Write-Host -BackgroundColor Yellow -ForegroundColor White "Copy to PrivateServ"
ssh root@8.222.171.134 "killall forLinux"
scp forLinux root@8.222.171.134:~
ssh root@8.222.171.134 "chmod +x ~/forLinux"
ssh root@8.222.171.134 "nohup ./forLinux > /dev/null 2>&1 &"

# windows vm-s19
Write-Host -BackgroundColor Yellow -ForegroundColor White "Copy to vm-s19"
ssh administrator@vm-s19 'Get-Process -Name forWin | Stop-Process -Force'
scp forWin.exe administrator@vm-s19:forWin.exe




# Update ALI OSS
# TODO::
