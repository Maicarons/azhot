# 构建系统文档

本文档详细说明了 azhot 项目的构建系统，包括本地构建和 CI/CD 集成。

## 概述

azhot 项目使用 CMake 作为主要的构建系统，支持：

- 本地开发构建
- 跨平台交叉编译
- CI/CD 集成
- 自动化发布流程

## CMake 构建系统

### 主要目标

- `build`: 构建当前平台的二进制文件
- `package`: 构建所有平台并打包为 zip 文件
- `release`: 构建并打包用于发布的版本
- `build-multiarch`: 构建多架构版本
- `test`: 运行测试
- `test-all`: 运行所有测试并生成覆盖率报告
- `dev`: 开发模式运行
- `run`: 构建并运行
- `azhot_clean`: 清理构建产物

### 跨平台构建目标

- `build-platform-linux`: 构建 Linux amd64 版本
- `build-platform-windows`: 构建 Windows amd64 版本
- `build-platform-darwin`: 构建 macOS amd64 版本
- `build-platform-linux-arm64`: 构建 Linux arm64 版本
- `build-platform-windows-arm64`: 构建 Windows arm64 版本

## CI/CD 集成

### GitHub Actions

项目包含以下 GitHub Actions 工作流：

- `release.yml`: 自动发布工作流，当推送标签时触发
  - 构建 Linux、Windows、macOS 三个平台的二进制文件
  - 为每个平台创建 zip 包
  - 自动发布到 GitHub Releases

### 构建流程

1. **测试阶段**: 运行单元测试确保代码质量
2. **构建阶段**: 为三个目标平台（Linux、Windows、macOS）构建二进制文件
3. **打包阶段**: 将每个平台的二进制文件打包为 zip 格式，包含必要的配置文件
4. **发布阶段**: 将打包文件上传到 GitHub Releases

## 本地构建

### 使用 CMake

```bash
# 基本构建
mkdir build && cd build
cmake ..
cmake --build . --target build

# 构建并打包用于发布
cmake --build . --target release
```

### 使用构建脚本

项目还提供独立的构建脚本：

- `scripts/build-release.sh` - Linux/macOS 构建脚本
- `scripts/build-release.bat` - Windows 构建脚本

```bash
# Linux/macOS
./scripts/build-release.sh [输出目录]

# Windows
./scripts\build-release.bat [输出目录]
```

## 构建脚本

### cmake_scripts/build_platform.cmake

辅助 CMake 脚本，用于构建特定平台：

```bash
cmake -P cmake_scripts/build_platform.cmake <platform> [output_dir] [package]
```

支持的平台：
- `linux`
- `windows`
- `darwin` 或 `macos`
- `linux-arm64`
- `windows-arm64`
- `linux-arm`
- `freebsd`

### cmake_scripts/build_all_platforms.cmake

构建所有平台并打包的脚本：

```bash
cmake -P cmake_scripts/build_all_platforms.cmake [output_dir]
```

## 发布流程

### 自动发布

1. 创建并推送 Git 标签（如 `v1.0.0`）
2. GitHub Actions 自动触发 `release.yml` 工作流
3. 工作流构建三个平台的二进制文件
4. 为每个平台创建 zip 包
5. 创建 GitHub Release 并上传包文件

### 包内容

每个平台的 zip 包包含：
- 二进制文件
- .env 配置文件（从 .env.example 复制）
- README.md
- LICENSE

## 配置

构建系统会自动检测操作系统并设置相应的环境变量：

- `GOOS`: 目标操作系统
- `GOARCH`: 目标架构
- `CGO_ENABLED=0`: 禁用 CGO 以确保可移植性
- 使用 `-ldflags="-s -w"` 以减小二进制文件大小

## 故障排除

### 常见问题

1. **缺少 Go 或 CMake**: 确保已安装 Go 1.18+ 和 CMake 3.12+
2. **交叉编译问题**: 确保设置了正确的环境变量
3. **权限问题**: 在 Linux/macOS 上确保脚本有执行权限

### 验证构建

构建完成后，可以验证生成的二进制文件：

```bash
file package/azhot_linux_amd64  # 检查文件类型
ls -la package/*.zip            # 检查包文件
```