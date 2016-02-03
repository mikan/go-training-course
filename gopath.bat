for /f "delims=" %%a in ('@cd') do setx GOPATH %%a
@echo off
rem in Command Prompt window:
rem > for /f "delims=" %a in ('@cd') do setx GOPATH %a
pause
