# Arm Board
Write-Host  -ForegroundColor White "Copy to ARM"
ssh openailab@192.168.3.146 "killall forArm"
scp forArm openailab@192.168.3.146:~
ssh openailab@192.168.3.146 "nohup ./forArm > /dev/null 2>&1 &"


# utools
Write-Host  -ForegroundColor White "Copy to utools"
ssh root@39.108.212.138 "killall forLinux"
scp forLinux root@39.108.212.138:~
ssh root@39.108.212.138 "nohup ./forLinux > /dev/null 2>&1 &"

# WebServ
Write-Host  -ForegroundColor White "Copy to WebServ"
ssh root@120.79.203.54 "killall forLinux"
scp forLinux root@120.79.203.54:~
ssh root@120.79.203.54 "nohup ./forLinux > /dev/null 2>&1 &"

# PrivateServ
Write-Host  -ForegroundColor White "Copy to PrivateServ"
ssh root@8.222.171.134 "killall forLinux"
scp forLinux root@8.222.171.134:~
ssh root@8.222.171.134 "chmod +x ~/forLinux"
ssh root@8.222.171.134 "nohup ./forLinux > /dev/null 2>&1 &"


# Update ALI OSS
# TODO::
$ss=New-PSSession -ComputerName vm-a,vm-s19 -Credential administrator

invoke-Command -Session $ss  -ScriptBlock {Set-NetFirewallProfile -All -Enabled False }

invoke-Command -Session $ss  -ScriptBlock {Stop-Process -Name forWin -ErrorAction SilentlyContinue}

Copy-Item C:\Users\sheng\ITAdminPlus\forWin.exe -Destination \\vm-a\c$\forWin.exe
Copy-Item C:\Users\sheng\ITAdminPlus\forWin.exe -Destination \\vm-s19\c$\forWin.exe

invoke-Command -Session $ss  -ScriptBlock {Start-Process c:\forWin.exe }
invoke-Command -Session $ss  -ScriptBlock {Get-Process -Name forWin}
