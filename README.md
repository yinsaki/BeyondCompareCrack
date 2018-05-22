# BeyondCompareCrack

## 打开命令行
### 查看当前用户的SID
`whoami /user`
用这里查到的SID替换下面的<SID>,每个人的SID不一样。

### 从注册表删除CacheId
```
reg delete "HKEY_USERS\<SID>\Software\Scooter Software\Beyond Compare 4" /v CacheId /f
```