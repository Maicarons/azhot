cmake_minimum_required(VERSION 3.12)
project(azhot)

# 设置项目变量
set(BINARY_NAME azhot)
set(BUILD_DIR build)
set(DIST_DIR dist)

# 检测操作系统
if(WIN32)
    set(TARGET_OS windows)
    set(EXECUTABLE_NAME ${BINARY_NAME}.exe)
    set(ZIP_EXTENSION .zip)
    set(PLATFORM_SCRIPT build_platform.bat)
else()
    if(CMAKE_SYSTEM_NAME STREQUAL "Linux")
        set(TARGET_OS linux)
    elseif(CMAKE_SYSTEM_NAME STREQUAL "Darwin")
        set(TARGET_OS darwin)
    else()
        set(TARGET_OS ${CMAKE_SYSTEM_NAME})
    endif()
    set(EXECUTABLE_NAME ${BINARY_NAME})
    set(ZIP_EXTENSION .tar.gz)
    set(PLATFORM_SCRIPT build_platform.sh)
endif()

# 添加自定义目标：构建当前平台
add_custom_target(build
    COMMAND go install github.com/swaggo/swag/cmd/swag@latest
    COMMAND swag init
    COMMAND ${CMAKE_COMMAND} -E make_directory ${BUILD_DIR}
    COMMAND go build -o ${BUILD_DIR}/${EXECUTABLE_NAME} .
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    COMMENT "Building for current platform (${CMAKE_SYSTEM_NAME})..."
)

# 跨平台构建命令函数
function(build_for_platform OS ARCH OUTPUT_NAME)
    if(WIN32)
        if(${OS} STREQUAL "darwin")
            # 对于 macOS 平台，禁用 CGO 以避免 -arch 参数问题，使用纯 Go SQLite 实现
            add_custom_target(build_${OS}_${ARCH}
                COMMAND ${CMAKE_COMMAND} -E make_directory ${BUILD_DIR}
                COMMAND cmd /c "set \"GOOS=${OS}\" && set \"GOARCH=${ARCH}\" && set \"CGO_ENABLED=0\" && go build -tags \"sqlite_omit_load_extension\" -o \"${BUILD_DIR}/${OUTPUT_NAME}\" ."
                WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
                COMMENT "Building for ${OS}/${ARCH}..."
            )
        else()
            add_custom_target(build_${OS}_${ARCH}
                COMMAND ${CMAKE_COMMAND} -E make_directory ${BUILD_DIR}
                COMMAND cmd /c "set \"GOOS=${OS}\" && set \"GOARCH=${ARCH}\" && set \"CGO_ENABLED=1\" && go build -o \"${BUILD_DIR}/${OUTPUT_NAME}\" ."
                WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
                COMMENT "Building for ${OS}/${ARCH}..."
            )
        endif()
    else()
        if(${OS} STREQUAL "darwin")
            # 对于 macOS 平台，禁用 CGO 以避免 -arch 参数问题，使用纯 Go SQLite 实现
            add_custom_target(build_${OS}_${ARCH}
                COMMAND ${CMAKE_COMMAND} -E make_directory ${BUILD_DIR}
                COMMAND env GOOS=${OS} GOARCH=${ARCH} CGO_ENABLED=0 go build -tags "sqlite_omit_load_extension" -o ${BUILD_DIR}/${OUTPUT_NAME} .
                WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
                COMMENT "Building for ${OS}/${ARCH}..."
            )
        else()
            add_custom_target(build_${OS}_${ARCH}
                COMMAND ${CMAKE_COMMAND} -E make_directory ${BUILD_DIR}
                COMMAND env GOOS=${OS} GOARCH=${ARCH} CGO_ENABLED=1 go build -o ${BUILD_DIR}/${OUTPUT_NAME} .
                WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
                COMMENT "Building for ${OS}/${ARCH}..."
            )
        endif()
    endif()
endfunction()

# 预定义平台构建
build_for_platform(linux amd64 ${BINARY_NAME}_linux)
build_for_platform(windows amd64 ${BINARY_NAME}_windows.exe)
build_for_platform(darwin amd64 ${BINARY_NAME}_darwin)

# 添加自定义目标：打包
add_custom_target(package
    # 先构建所有平台
    DEPENDS build_linux_amd64 build_windows_amd64 build_darwin_amd64
    COMMAND ${CMAKE_COMMAND} -E make_directory ${DIST_DIR}/linux-amd64
    COMMAND ${CMAKE_COMMAND} -E make_directory ${DIST_DIR}/windows-amd64
    COMMAND ${CMAKE_COMMAND} -E make_directory ${DIST_DIR}/darwin-amd64
    COMMAND ${CMAKE_COMMAND} -E copy ${BUILD_DIR}/${BINARY_NAME}_linux ${DIST_DIR}/linux-amd64/${BINARY_NAME}
    COMMAND ${CMAKE_COMMAND} -E copy ${BUILD_DIR}/${BINARY_NAME}_windows.exe ${DIST_DIR}/windows-amd64/${BINARY_NAME}.exe
    COMMAND ${CMAKE_COMMAND} -E copy ${BUILD_DIR}/${BINARY_NAME}_darwin ${DIST_DIR}/darwin-amd64/${BINARY_NAME}
    COMMAND ${CMAKE_COMMAND} -E copy .env.example ${DIST_DIR}/linux-amd64/.env
    COMMAND ${CMAKE_COMMAND} -E copy .env.example ${DIST_DIR}/windows-amd64/.env
    COMMAND ${CMAKE_COMMAND} -E copy .env.example ${DIST_DIR}/darwin-amd64/.env
    COMMAND ${CMAKE_COMMAND} -E copy README.md ${DIST_DIR}/linux-amd64/
    COMMAND ${CMAKE_COMMAND} -E copy README.md ${DIST_DIR}/windows-amd64/
    COMMAND ${CMAKE_COMMAND} -E copy README.md ${DIST_DIR}/darwin-amd64/
    COMMAND ${CMAKE_COMMAND} -E copy LICENSE ${DIST_DIR}/linux-amd64/
    COMMAND ${CMAKE_COMMAND} -E copy LICENSE ${DIST_DIR}/windows-amd64/
    COMMAND ${CMAKE_COMMAND} -E copy LICENSE ${DIST_DIR}/darwin-amd64/
    # 创建zip压缩包
    COMMAND ${CMAKE_COMMAND} -E chdir ${DIST_DIR}/linux-amd64/ ${CMAKE_COMMAND} -E tar "cf" "../azhot_linux_amd64.zip" "--format=zip" "."
    COMMAND ${CMAKE_COMMAND} -E chdir ${DIST_DIR}/windows-amd64/ ${CMAKE_COMMAND} -E tar "cf" "../azhot_windows_amd64.zip" "--format=zip" "."
    COMMAND ${CMAKE_COMMAND} -E chdir ${DIST_DIR}/darwin-amd64/ ${CMAKE_COMMAND} -E tar "cf" "../azhot_darwin_amd64.zip" "--format=zip" "."
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    COMMENT "Packaging for all platforms..."
)

# 添加自定义目标：清理构建产物
add_custom_target(azhot_clean
    COMMAND ${CMAKE_COMMAND} -E remove_directory ${DIST_DIR}
    COMMAND go clean -cache
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    COMMENT "Cleaning build artifacts..."
)

# 添加自定义目标：运行
add_custom_target(run
    COMMAND ${CMAKE_COMMAND} -E make_directory ${BUILD_DIR}
    COMMAND go build -o ${BUILD_DIR}/${EXECUTABLE_NAME} .
    COMMAND ${CMAKE_COMMAND} -E chdir ${BUILD_DIR} ${EXECUTABLE_NAME}
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    COMMENT "Running application..."
)

# 添加自定义目标：开发模式
add_custom_target(dev
    COMMAND go run main.go
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    COMMENT "Starting development mode..."
)

# 添加自定义目标：测试
add_custom_target(test
    COMMAND go test -v ./...
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    COMMENT "Running tests..."
)

# 添加自定义目标：测试所有
if(WIN32)
    add_custom_target(test-all
        COMMAND cmd /c "chcp 65001 >nul && go test -v -coverprofile=coverage.out ./..."
        WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
        COMMENT "Running all tests..."
    )
else()
    add_custom_target(test-all
        COMMAND go test -v -coverprofile=coverage.out ./...
        WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
        COMMENT "Running all tests..."
    )
endif()

# 添加自定义目标：格式化
add_custom_target(fmt
    COMMAND go fmt ./...
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    COMMENT "Formatting code..."
)

# 添加自定义目标：vet
add_custom_target(vet
    COMMAND go vet ./...
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    COMMENT "Running go vet..."
)

# 添加自定义目标：依赖整理
add_custom_target(tidy
    COMMAND go mod tidy
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    COMMENT "Tidying go.mod..."
)

# 添加自定义目标：静态检查
add_custom_target(staticcheck
    COMMAND go install honnef.co/go/tools/cmd/staticcheck@latest
    COMMAND staticcheck ./...
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    COMMENT "Running static analysis..."
)

# 添加自定义目标：构建CI版本（不生成swagger文档）
add_custom_target(build-ci
    COMMAND ${CMAKE_COMMAND} -E make_directory ${BUILD_DIR}
    COMMAND go build -o ${BUILD_DIR}/${EXECUTABLE_NAME} .
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    COMMENT "Building for CI environment..."
)

# 添加自定义目标：跨平台构建
add_custom_target(build-platform
    COMMAND ${CMAKE_COMMAND} -E echo "Usage: cmake --build . --target build-platform-linux, build-platform-windows, or build-platform-darwin"
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    COMMENT "Cross-platform build target."
)

# 为跨平台构建添加特定目标
if(WIN32)
    add_custom_target(build-platform-linux
        COMMAND ${CMAKE_COMMAND} -E make_directory ${BUILD_DIR}
        COMMAND cmd /c "set \"GOOS=linux\" && set \"GOARCH=amd64\" && set \"CGO_ENABLED=0\" && go build -o \"${BUILD_DIR}/${BINARY_NAME}_linux\" ."
        WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
        COMMENT "Building for Linux amd64..."
    )

    add_custom_target(build-platform-windows
        COMMAND ${CMAKE_COMMAND} -E make_directory ${BUILD_DIR}
        COMMAND cmd /c "set \"GOOS=windows\" && set \"GOARCH=amd64\" && set \"CGO_ENABLED=0\" && go build -o \"${BUILD_DIR}/${BINARY_NAME}_windows.exe\" ."
        WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
        COMMENT "Building for Windows amd64..."
    )

    add_custom_target(build-platform-darwin
        COMMAND ${CMAKE_COMMAND} -E make_directory ${BUILD_DIR}
        COMMAND cmd /c "set \"GOOS=darwin\" && set \"GOARCH=amd64\" && set \"CGO_ENABLED=0\" && go build -o \"${BUILD_DIR}/${BINARY_NAME}_darwin\" ."
        WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
        COMMENT "Building for Darwin (macOS) amd64..."
    )

    add_custom_target(build-platform-linux-arm64
        COMMAND ${CMAKE_COMMAND} -E make_directory ${BUILD_DIR}
        COMMAND cmd /c "set \"GOOS=linux\" && set \"GOARCH=arm64\" && set \"CGO_ENABLED=0\" && go build -o \"${BUILD_DIR}/${BINARY_NAME}_linux_arm64\" ."
        WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
        COMMENT "Building for Linux arm64..."
    )

    add_custom_target(build-platform-windows-arm64
        COMMAND ${CMAKE_COMMAND} -E make_directory ${BUILD_DIR}
        COMMAND cmd /c "set \"GOOS=windows\" && set \"GOARCH=arm64\" && set \"CGO_ENABLED=0\" && go build -o \"${BUILD_DIR}/${BINARY_NAME}_windows_arm64.exe\" ."
        WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
        COMMENT "Building for Windows arm64..."
    )
else()
    add_custom_target(build-platform-linux
        COMMAND ${CMAKE_COMMAND} -E make_directory ${BUILD_DIR}
        COMMAND env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${BUILD_DIR}/${BINARY_NAME}_linux .
        WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
        COMMENT "Building for Linux amd64..."
    )

    add_custom_target(build-platform-windows
        COMMAND ${CMAKE_COMMAND} -E make_directory ${BUILD_DIR}
        COMMAND env GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o ${BUILD_DIR}/${BINARY_NAME}_windows.exe .
        WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
        COMMENT "Building for Windows amd64..."
    )

    add_custom_target(build-platform-darwin
        COMMAND ${CMAKE_COMMAND} -E make_directory ${BUILD_DIR}
        COMMAND env GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o ${BUILD_DIR}/${BINARY_NAME}_darwin .
        WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
        COMMENT "Building for Darwin (macOS) amd64..."
    )

    add_custom_target(build-platform-linux-arm64
        COMMAND ${CMAKE_COMMAND} -E make_directory ${BUILD_DIR}
        COMMAND env GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o ${BUILD_DIR}/${BINARY_NAME}_linux_arm64 .
        WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
        COMMENT "Building for Linux arm64..."
    )

    add_custom_target(build-platform-windows-arm64
        COMMAND ${CMAKE_COMMAND} -E make_directory ${BUILD_DIR}
        COMMAND env GOOS=windows GOARCH=arm64 CGO_ENABLED=0 go build -o ${BUILD_DIR}/${BINARY_NAME}_windows_arm64.exe .
        WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
        COMMENT "Building for Windows arm64..."
    )
endif()
