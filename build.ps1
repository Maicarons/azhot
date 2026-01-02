# PowerShell 构建脚本 for azhot

param(
    [string]$Command = "build",
    [switch]$Help
)

if ($Help) {
    Write-Host "azhot PowerShell Build Script"
    Write-Host ""
    Write-Host "Usage:"
    Write-Host "  .\build.ps1 [command]"
    Write-Host ""
    Write-Host "Commands:"
    Write-Host "  build       Build the application"
    Write-Host "  run         Build and run the application"
    Write-Host "  dev         Run in development mode"
    Write-Host "  test        Run tests"
    Write-Host "  fmt         Format code"
    Write-Host "  tidy        Tidy go.mod"
    Write-Host "  clean       Clean build artifacts"
    Write-Host "  package     Package application with necessary files"
    Write-Host "  help        Show this help"
    exit 0
}

switch ($Command) {
    "build" {
        Write-Host "Building hotsearch_api..." -ForegroundColor Green
        if (!(Test-Path "build")) {
            New-Item -ItemType Directory -Path "build" | Out-Null
        }
        go build -o build/hotsearch_api.exe .
        if ($LASTEXITCODE -eq 0) {
            Write-Host "Build successful!" -ForegroundColor Green
        } else {
            Write-Host "Build failed!" -ForegroundColor Red
            exit 1
        }
    }
    "run" {
        Write-Host "Running azhot..." -ForegroundColor Green
        .\build.ps1 build
        .\build\hotsearch_api.exe
    }
    "dev" {
        Write-Host "Starting development mode..." -ForegroundColor Green
        go run main.go
    }
    "test" {
        Write-Host "Running tests..." -ForegroundColor Green
        go test -v ./...
    }
    "fmt" {
        Write-Host "Formatting code..." -ForegroundColor Green
        go fmt ./...
    }
    "tidy" {
        Write-Host "Tidying go.mod..." -ForegroundColor Green
        go mod tidy
    }
    "clean" {
        Write-Host "Cleaning build artifacts..." -ForegroundColor Green
        if (Test-Path "build") {
            Remove-Item -Recurse -Force build
        }
        if (Test-Path "dist") {
            Remove-Item -Recurse -Force dist
        }
        go clean -cache
    }
    "package" {
        Write-Host "Packaging application..." -ForegroundColor Green
        .\build.ps1 build
        
        if (!(Test-Path "dist")) {
            New-Item -ItemType Directory -Path "dist" | Out-Null
        }
        
        # 创建发布目录
        $packageName = "hotsearch_api_windows"
        $packageDir = "dist\$packageName"
        if (Test-Path $packageDir) {
            Remove-Item -Recurse -Force $packageDir
        }
        New-Item -ItemType Directory -Path $packageDir | Out-Null
        
        # 复制构建产物和必备文件
        Copy-Item "build\hotsearch_api.exe" $packageDir
        if (Test-Path ".env.example") { Copy-Item ".env.example" $packageDir }
        if (Test-Path "hot_search.db") { Copy-Item "hot_search.db" $packageDir }
        if (Test-Path "README.md") { Copy-Item "README.md" $packageDir }
        if (Test-Path "README_zh.md") { Copy-Item "README_zh.md" $packageDir }
        if (Test-Path "LICENSE") { Copy-Item "LICENSE" $packageDir }
        
        # 创建zip包
        $zipPath = "dist\hotsearch_api_current_platform.zip"
        if (Test-Path $zipPath) { Remove-Item $zipPath }
        Compress-Archive -Path "$packageDir\*" -DestinationPath $zipPath -Force
        
        Write-Host "Packaging completed. Archive created: $zipPath" -ForegroundColor Green
    }
    default {
        Write-Host "Unknown command: $Command" -ForegroundColor Red
        Write-Host "Use .\build.ps1 -Help for usage information"
    }
}