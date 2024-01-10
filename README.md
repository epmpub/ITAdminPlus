
# 测试流程

### compile and upload bin and md5.json to OSS .

### call powershell script
```powershell
   {
   	"type": "powershell_script",
   	"command":"curl -s https://it2u.oss-cn-shenzhen.aliyuncs.com/scripts/test.ps1 | powershell -NoLogo"
   }

```

### call bash script call

```bash
   {
   	"type": "powershell_script",
   	"command":"curl -s https://it2u.oss-cn-shenzhen.aliyuncs.com/scripts/test.ps1 | powershell -NoLogo"
   }

```

