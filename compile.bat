@echo off
setlocal

REM Set environment variables for cross-compilation
set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0

REM Set output directory
set OUTPUT_DIR=D:\GoPath\user_growth\dockerfile

REM Create output directory if it doesn't exist
if not exist %OUTPUT_DIR% mkdir %OUTPUT_DIR%

REM Compile client
echo Compiling client...
go build -o %OUTPUT_DIR%\client ./client

REM Compile server
echo Compiling server...
go build -o %OUTPUT_DIR%\server ./server

REM Compile gin
echo Compiling gin...
go build -o %OUTPUT_DIR%\gin_app ./maingin

REM Compile grpcurl_linux
echo Compiling grpcurl_linux ...
go build -o %OUTPUT_DIR%\grpcurl_linux github.com/fullstorydev/grpcurl/cmd/grpcurl

echo Compilation complete. Binaries are in %OUTPUT_DIR%