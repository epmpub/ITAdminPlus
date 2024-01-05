# Arm Dev
Get-Content $env:USERPROFILE\.ssh\id_rsa.pub | ssh openailab@192.168.3.146 "cat >> .ssh/authorized_keys"
Get-Content $env:USERPROFILE\.ssh\id_rsa.pub | ssh openailab@192.168.3.172 "cat >> .ssh/authorized_keys"


# utools
Get-Content $env:USERPROFILE\.ssh\id_rsa.pub | ssh root@39.108.212.138 "cat >> .ssh/authorized_keys"

# WebServ
Get-Content $env:USERPROFILE\.ssh\id_rsa.pub | ssh root@120.79.203.54 "cat >> .ssh/authorized_keys"

# VM-A (win10) NOT Work.

$authorizedKey = Get-Content -Path $env:USERPROFILE\.ssh\id_ed25519.pub
$remotePowershell = "powershell Add-Content -Force -Path $env:ProgramData\ssh\administrators_authorized_keys -Value '$authorizedKey';icacls.exe ""$env:ProgramData\ssh\administrators_authorized_keys"" /inheritance:r /grant ""Administrators:F"" /grant ""SYSTEM:F"""
ssh clouduser@vm-s19 $remotePowershell
ssh clouduser@vm-s22 $remotePowershell

