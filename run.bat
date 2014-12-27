@echo off

set FILE_NAME=%1
set FILE_PATH=%2
set PROJECT_PATH=%3

set GOPATH=%PROJECT_PATH%
echo "installing..."
go install
echo "runing..."
go run %FILE_NAME%