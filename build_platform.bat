@echo off
REM 跨平台构建批处理脚本

if "%1"=="" (
    echo Usage: %0 ^<platform^>
    echo Supported platforms: linux, windows, darwin/macos, linux-arm64, windows-arm64
    exit /b 1
)

set PLATFORM=%1

if "%PLATFORM%"=="linux" (
    set GOOS=linux
    set GOARCH=amd64
    set CGO_ENABLED=0
    go build -o build\azhot_linux.exe .
) else if "%PLATFORM%"=="windows" (
    set GOOS=windows
    set GOARCH=amd64
    set CGO_ENABLED=0
    go build -o build\azhot_windows.exe .
) else if "%PLATFORM%"=="darwin" (
    set GOOS=darwin
    set GOARCH=amd64
    set CGO_ENABLED=0
    go build -o build\azhot_darwin .
) else if "%PLATFORM%"=="macos" (
    set GOOS=darwin
    set GOARCH=amd64
    set CGO_ENABLED=0
    go build -o build\azhot_darwin .
) else if "%PLATFORM%"=="linux-arm64" (
    set GOOS=linux
    set GOARCH=arm64
    set CGO_ENABLED=0
    go build -o build\azhot_linux_arm64 .
) else if "%PLATFORM%"=="windows-arm64" (
    set GOOS=windows
    set GOARCH=arm64
    set CGO_ENABLED=0
    go build -o build\azhot_windows_arm64.exe .
) else if "%PLATFORM%"=="linux-arm" (
    set GOOS=linux
    set GOARCH=arm
    set CGO_ENABLED=0
    go build -o build\azhot_linux_arm .
) else if "%PLATFORM%"=="freebsd" (
    set GOOS=freebsd
    set GOARCH=amd64
    set CGO_ENABLED=0
    go build -o build\azhot_freebsd .
) else (
    echo Unsupported platform: %PLATFORM%
    echo Supported platforms: linux, windows, darwin/macos, linux-arm64, windows-arm64, linux-arm, freebsd
    exit /b 1
)

if %ERRORLEVEL% EQU 0 (
    echo Successfully built for %PLATFORM%
) else (
    echo Build failed for %PLATFORM%
    exit /b 1
)