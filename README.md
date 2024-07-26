# ntweakpwdscan
This is a tool to find windows weak password by ntlm hash comparison.

Usage:
1.dump windows ntlm hash to file hash.txt with **[mimikatz](https://github.com/gentilkiwi/mimikatz)**.
```
mimikatz.exe  "privilege::debug"   "lsadump::lsa /patch"   exit>hash.txt
```
2.compare the exported ntlm hash file(hash.txt) and your weak password dict(password.txt) with ntweakpwdscan tool,and generate result to file accounts_xxxxx.xlsx default.
```
ntweakpwdscan.exe -f hash.txt -pass password.txt
```
***
## One-click completion

You can put the password.txt、ntweakpwdscan.exe、checkout.bat file into the directory of **[mimikatz.exe](https://github.com/gentilkiwi/mimikatz/releases)**,then right-click checkout.bat and run as an administrator to automatically complete the above steps.