#!/bin/bash

# 构建发布版本的脚本
# 用于在CI环境中构建所有平台并打包

set -e  # 遇到错误时退出

echo "开始构建发布版本..."

# 检查是否安装了必要的工具
if ! command -v go &> /dev/null; then
    echo "错误: 未找到 go 命令，请安装 Go"
    exit 1
fi

if ! command -v cmake &> /dev/null; then
    echo "错误: 未找到 cmake 命令，请安装 CMake"
    exit 1
fi

# 设置输出目录
OUTPUT_DIR=${1:-"package"}
mkdir -p "$OUTPUT_DIR"

# 设置环境变量
export CGO_ENABLED=0

echo "构建 Linux amd64 版本..."
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o "$OUTPUT_DIR/azhot_linux_amd64" .

echo "构建 Windows amd64 版本..."
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o "$OUTPUT_DIR/azhot_windows_amd64.exe" .

echo "构建 Darwin (macOS) amd64 版本..."
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o "$OUTPUT_DIR/azhot_darwin_amd64" .

echo "创建 Linux 包..."
mkdir -p "$OUTPUT_DIR/temp_linux"
cp "$OUTPUT_DIR/azhot_linux_amd64" "$OUTPUT_DIR/temp_linux/"
cp -f .env.example "$OUTPUT_DIR/temp_linux/.env" 2>/dev/null || echo ".env.example not found, skipping..."
cp -f README.md "$OUTPUT_DIR/temp_linux/" 2>/dev/null || echo "README.md not found, skipping..."
cp -f LICENSE "$OUTPUT_DIR/temp_linux/" 2>/dev/null || echo "LICENSE not found, skipping..."

cd "$OUTPUT_DIR/temp_linux" && zip -r "../azhot_linux_amd64.zip" . && cd - > /dev/null

echo "创建 Windows 包..."
mkdir -p "$OUTPUT_DIR/temp_windows"
cp "$OUTPUT_DIR/azhot_windows_amd64.exe" "$OUTPUT_DIR/temp_windows/"
cp -f .env.example "$OUTPUT_DIR/temp_windows/.env" 2>/dev/null || echo ".env.example not found, skipping..."
cp -f README.md "$OUTPUT_DIR/temp_windows/" 2>/dev/null || echo "README.md not found, skipping..."
cp -f LICENSE "$OUTPUT_DIR/temp_windows/" 2>/dev/null || echo "LICENSE not found, skipping..."

cd "$OUTPUT_DIR/temp_windows" && zip -r "../azhot_windows_amd64.zip" . && cd - > /dev/null

echo "创建 Darwin (macOS) 包..."
mkdir -p "$OUTPUT_DIR/temp_darwin"
cp "$OUTPUT_DIR/azhot_darwin_amd64" "$OUTPUT_DIR/temp_darwin/"
cp -f .env.example "$OUTPUT_DIR/temp_darwin/.env" 2>/dev/null || echo ".env.example not found, skipping..."
cp -f README.md "$OUTPUT_DIR/temp_darwin/" 2>/dev/null || echo "README.md not found, skipping..."
cp -f LICENSE "$OUTPUT_DIR/temp_darwin/" 2>/dev/null || echo "LICENSE not found, skipping..."

cd "$OUTPUT_DIR/temp_darwin" && zip -r "../azhot_darwin_amd64.zip" . && cd - > /dev/null

# 清理临时目录
rm -rf "$OUTPUT_DIR/temp_linux" "$OUTPUT_DIR/temp_windows" "$OUTPUT_DIR/temp_darwin"

echo "构建和打包完成！输出文件在 $OUTPUT_DIR/ 目录中："
ls -la "$OUTPUT_DIR"/*.zip

echo "验证二进制文件..."
file "$OUTPUT_DIR/azhot_linux_amd64" 2>/dev/null || echo "file command not available"
file "$OUTPUT_DIR/azhot_darwin_amd64" 2>/dev/null || echo "file command not available"