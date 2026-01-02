# azhot Makefile
# 用于构建、测试和运行热搜API服务

# 变量定义
BINARY_NAME=hotsearch_api
MAIN_PATH=.
BUILD_DIR=build
DIST_DIR=dist
VERSION=1.0.0

# 检测操作系统
ifeq ($(OS),Windows_NT)
    TARGET_OS := windows
else
    TARGET_OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')
endif

# 默认目标 - 编译当前平台
.PHONY: all
all: build

# 生成API文档
.PHONY: swagger
swagger:
	@echo "Generating API documentation..."
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init

# 构建应用前先生成API文档 (当前平台)
.PHONY: build
build: swagger
	@echo "Building for current platform ($(TARGET_OS))..."
	@mkdir -p $(BUILD_DIR)
ifeq ($(TARGET_OS),windows)
	go build -o $(BUILD_DIR)/$(BINARY_NAME).exe $(MAIN_PATH)
else
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
endif

# 清理构建产物
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR) $(DIST_DIR)
	go clean -cache

# 运行应用
.PHONY: run
run: build
	@echo "Starting server..."
ifeq ($(TARGET_OS),windows)
	./$(BUILD_DIR)/$(BINARY_NAME).exe
else
	./$(BUILD_DIR)/$(BINARY_NAME)
endif

# 开发模式运行
.PHONY: dev
dev:
	@echo "Starting development mode..."
	go run main.go

# 测试
.PHONY: test
test:
	@echo "Running tests..."
	go test -v ./...

# 测试（覆盖所有包）
.PHONY: test-all
test-all:
	@echo "Running all tests..."
	go test -v -coverprofile=coverage.out ./...

# 格式化代码
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	go fmt ./...

# 检查代码错误
.PHONY: vet
vet:
	@echo "Checking code with go vet..."
	go vet ./...

# 更新依赖
.PHONY: tidy
tidy:
	@echo "Tidying go.mod..."
	go mod tidy

# 交叉编译 - Linux
.PHONY: build-linux
build-linux:
	@echo "Building for Linux..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)_linux $(MAIN_PATH)

# 交叉编译 - Windows
.PHONY: build-windows
build-windows:
	@echo "Building for Windows..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)_windows.exe $(MAIN_PATH)

# 打包当前平台
.PHONY: package-current
package-current: build
	@echo "Packaging for current platform..."
	@mkdir -p $(DIST_DIR)/current-platform
	cp -r $(BUILD_DIR)/$(BINARY_NAME)* $(DIST_DIR)/current-platform/
	@echo "Current platform package created in $(DIST_DIR)/current-platform directory."

# 打包所有平台
.PHONY: package
package: build-linux build-windows
	@echo "Packaging for all platforms..."
	@mkdir -p $(DIST_DIR)/linux-amd64 $(DIST_DIR)/windows-amd64
	cp $(BUILD_DIR)/$(BINARY_NAME)_linux $(DIST_DIR)/linux-amd64/
	cp $(BUILD_DIR)/$(BINARY_NAME)_windows.exe $(DIST_DIR)/windows-amd64/
	@echo "Packages created in $(DIST_DIR)/ directory."

# 帮助信息
.PHONY: help
help:
	@echo "azhot Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make build              Build for current platform"
	@echo "  make run                Build and run for current platform"
	@echo "  make dev                Run in development mode"
	@echo "  make test               Run tests"
	@echo "  make test-all           Run all tests with coverage"
	@echo "  make fmt                Format code with go fmt"
	@echo "  make vet                Check code with go vet"
	@echo "  make tidy               Tidy go.mod"
	@echo "  make build-linux        Build for Linux (cross-compilation)"
	@echo "  make build-windows      Build for Windows (cross-compilation)"
	@echo "  make package            Package for all platforms"
	@echo "  make package-current    Package for current platform"
	@echo "  make clean              Clean build artifacts"
	@echo "  make help               Show this help"