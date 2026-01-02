#!/bin/bash
# 跨平台构建脚本

if [ $# -eq 0 ]; then
    echo "Usage: $0 <platform>"
    echo "Supported platforms: linux, windows, darwin/macos, linux-arm64, windows-arm64"
    exit 1
fi

PLATFORM=$1

case $PLATFORM in
    "linux")
        GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/azhot_linux .
        ;;
    "windows")
        GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o build/azhot_windows.exe .
        ;;
    "darwin"|"macos")
        GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o build/azhot_darwin .
        ;;
    "linux-arm64")
        GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o build/azhot_linux_arm64 .
        ;;
    "windows-arm64")
        GOOS=windows GOARCH=arm64 CGO_ENABLED=0 go build -o build/azhot_windows_arm64.exe .
        ;;
    "linux-arm")
        GOOS=linux GOARCH=arm CGO_ENABLED=0 go build -o build/azhot_linux_arm .
        ;;
    "freebsd")
        GOOS=freebsd GOARCH=amd64 CGO_ENABLED=0 go build -o build/azhot_freebsd .
        ;;
    *)
        echo "Unsupported platform: $PLATFORM"
        echo "Supported platforms: linux, windows, darwin/macos, linux-arm64, windows-arm64, linux-arm, freebsd"
        exit 1
        ;;
esac

if [ $? -eq 0 ]; then
    echo "Successfully built for $PLATFORM"
else
    echo "Build failed for $PLATFORM"
    exit 1
fi