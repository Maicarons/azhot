@echo off
REM 构建发布版本的批处理脚本
REM 用于在Windows CI环境中构建所有平台并打包

echo 开始构建发布版本...

REM 检查是否安装了必要的工具
go version >nul 2>&1
if errorlevel 1 (
    echo 错误: 未找到 go 命令，请安装 Go
    exit /b 1
)

cmake --version >nul 2>&1
if errorlevel 1 (
    echo 错误: 未找到 cmake 命令，请安装 CMake
    exit /b 1
)

REM 设置输出目录
set "OUTPUT_DIR=%1"
if "%OUTPUT_DIR%"=="" set "OUTPUT_DIR=package"
if not exist "%OUTPUT_DIR%" mkdir "%OUTPUT_DIR%"

REM 设置环境变量
set "CGO_ENABLED=0"

echo 构建 Linux amd64 版本...
set "GOOS=linux"
set "GOARCH=amd64"
go build -ldflags="-s -w" -o "%OUTPUT_DIR%\azhot_linux_amd64" .
if errorlevel 1 (
    echo 构建 Linux amd64 版本失败
    exit /b 1
)

echo 构建 Windows amd64 版本...
set "GOOS=windows"
set "GOARCH=amd64"
go build -ldflags="-s -w" -o "%OUTPUT_DIR%\azhot_windows_amd64.exe" .
if errorlevel 1 (
    echo 构建 Windows amd64 版本失败
    exit /b 1
)

echo 构建 Darwin (macOS) amd64 版本...
set "GOOS=darwin"
set "GOARCH=amd64"
go build -ldflags="-s -w" -o "%OUTPUT_DIR%\azhot_darwin_amd64" .
if errorlevel 1 (
    echo 构建 Darwin amd64 版本失败
    exit /b 1
)

REM 创建临时目录用于打包
echo 创建 Linux 包...
if exist "%OUTPUT_DIR%\temp_linux" rmdir /s /q "%OUTPUT_DIR%\temp_linux"
mkdir "%OUTPUT_DIR%\temp_linux"
copy "%OUTPUT_DIR%\azhot_linux_amd64" "%OUTPUT_DIR%\temp_linux\" >nul
if exist ".env.example" copy ".env.example" "%OUTPUT_DIR%\temp_linux\.env" >nul
if exist "README.md" copy "README.md" "%OUTPUT_DIR%\temp_linux\" >nul
if exist "LICENSE" copy "LICENSE" "%OUTPUT_DIR%\temp_linux\" >nul

REM 使用PowerShell创建zip文件（因为Windows自带的压缩工具）
powershell -Command "Compress-Archive -Path '%OUTPUT_DIR%\temp_linux\*' -DestinationPath '%OUTPUT_DIR%\azhot_linux_amd64.zip' -Force"

echo 创建 Windows 包...
if exist "%OUTPUT_DIR%\temp_windows" rmdir /s /q "%OUTPUT_DIR%\temp_windows"
mkdir "%OUTPUT_DIR%\temp_windows"
copy "%OUTPUT_DIR%\azhot_windows_amd64.exe" "%OUTPUT_DIR%\temp_windows\" >nul
if exist ".env.example" copy ".env.example" "%OUTPUT_DIR%\temp_windows\.env" >nul
if exist "README.md" copy "README.md" "%OUTPUT_DIR%\temp_windows\" >nul
if exist "LICENSE" copy "LICENSE" "%OUTPUT_DIR%\temp_windows\" >nul

powershell -Command "Compress-Archive -Path '%OUTPUT_DIR%\temp_windows\*' -DestinationPath '%OUTPUT_DIR%\azhot_windows_amd64.zip' -Force"

echo 创建 Darwin (macOS) 包...
if exist "%OUTPUT_DIR%\temp_darwin" rmdir /s /q "%OUTPUT_DIR%\temp_darwin"
mkdir "%OUTPUT_DIR%\temp_darwin"
copy "%OUTPUT_DIR%\azhot_darwin_amd64" "%OUTPUT_DIR%\temp_darwin\" >nul
if exist ".env.example" copy ".env.example" "%OUTPUT_DIR%\temp_darwin\.env" >nul
if exist "README.md" copy "README.md" "%OUTPUT_DIR%\temp_darwin\" >nul
if exist "LICENSE" copy "LICENSE" "%OUTPUT_DIR%\temp_darwin\" >nul

powershell -Command "Compress-Archive -Path '%OUTPUT_DIR%\temp_darwin\*' -DestinationPath '%OUTPUT_DIR%\azhot_darwin_amd64.zip' -Force"

REM 清理临时目录
if exist "%OUTPUT_DIR%\temp_linux" rmdir /s /q "%OUTPUT_DIR%\temp_linux"
if exist "%OUTPUT_DIR%\temp_windows" rmdir /s /q "%OUTPUT_DIR%\temp_windows"
if exist "%OUTPUT_DIR%\temp_darwin" rmdir /s /q "%OUTPUT_DIR%\temp_darwin"

echo 构建和打包完成！输出文件在 %OUTPUT_DIR% 目录中：
dir "%OUTPUT_DIR%\*.zip"

echo 构建发布版本完成。