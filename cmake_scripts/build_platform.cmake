#!/usr/bin/env cmake -P
# 这是一个辅助CMake脚本来处理跨平台构建
# 使用方法: cmake -P cmake_scripts/build_platform.cmake <platform>

cmake_minimum_required(VERSION 3.12)

# 获取平台参数
if(NOT CMAKE_ARGV3)
    message(FATAL_ERROR "No platform specified. Usage: cmake -P cmake_scripts/build_platform.cmake <platform>")
endif()

set(PLATFORM_ARG ${CMAKE_ARGV3})

# 定义平台映射
if(${PLATFORM_ARG} STREQUAL "linux")
    set(GOOS "linux")
    set(GOARCH "amd64")
    set(OUTPUT_NAME "azhot_linux")
elseif(${PLATFORM_ARG} STREQUAL "windows")
    set(GOOS "windows")
    set(GOARCH "amd64")
    set(OUTPUT_NAME "azhot_windows.exe")
elseif(${PLATFORM_ARG} STREQUAL "darwin" OR ${PLATFORM_ARG} STREQUAL "macos")
    set(GOOS "darwin")
    set(GOARCH "amd64")
    set(OUTPUT_NAME "azhot_darwin")
elseif(${PLATFORM_ARG} STREQUAL "linux-arm64")
    set(GOOS "linux")
    set(GOARCH "arm64")
    set(OUTPUT_NAME "azhot_linux_arm64")
elseif(${PLATFORM_ARG} STREQUAL "windows-arm64")
    set(GOOS "windows")
    set(GOARCH "arm64")
    set(OUTPUT_NAME "azhot_windows_arm64.exe")
elseif(${PLATFORM_ARG} STREQUAL "linux-arm")
    set(GOOS "linux")
    set(GOARCH "arm")
    set(OUTPUT_NAME "azhot_linux_arm")
elseif(${PLATFORM_ARG} STREQUAL "freebsd")
    set(GOOS "freebsd")
    set(GOARCH "amd64")
    set(OUTPUT_NAME "azhot_freebsd")
else()
    message(FATAL_ERROR "Unsupported platform: ${PLATFORM_ARG}. Supported platforms: linux, windows, darwin/macos, linux-arm64, windows-arm64, linux-arm, freebsd")
endif()

# 创建构建目录
file(MAKE_DIRECTORY build)

# 执行交叉编译
find_program(GO_CMD go)
if(NOT GO_CMD)
    message(FATAL_ERROR "go command not found. Please install Go.")
endif()

# 设置环境变量并执行构建
set(ENV_COMMAND)
if(WIN32)
    set(ENV_COMMAND cmd /c "set GOOS=${GOOS} && set GOARCH=${GOARCH} && set CGO_ENABLED=0 && go build -o build/${OUTPUT_NAME} .")
else()
    set(ENV_COMMAND env GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=0 go build -o build/${OUTPUT_NAME} .)
endif()

execute_process(
    COMMAND ${ENV_COMMAND}
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    RESULT_VARIABLE BUILD_RESULT
    OUTPUT_VARIABLE BUILD_OUTPUT
    ERROR_VARIABLE BUILD_ERROR
)

if(NOT BUILD_RESULT EQUAL 0)
    message(FATAL_ERROR "Build failed for platform ${PLATFORM_ARG}\nOutput: ${BUILD_OUTPUT}\nError: ${BUILD_ERROR}")
endif()

message(STATUS "Successfully built for ${PLATFORM_ARG} -> build/${OUTPUT_NAME}")