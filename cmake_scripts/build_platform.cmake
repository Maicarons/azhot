#!/usr/bin/env cmake -P
# 这是一个辅助CMake脚本来处理跨平台构建
# 使用方法: cmake -P cmake_scripts/build_platform.cmake <platform> [output_dir]

cmake_minimum_required(VERSION 3.12)

# 获取平台参数
if(NOT CMAKE_ARGV3)
    message(FATAL_ERROR "No platform specified. Usage: cmake -P cmake_scripts/build_platform.cmake <platform> [output_dir]")
endif()

set(PLATFORM_ARG ${CMAKE_ARGV3})

# 获取输出目录参数（可选），默认为build
if(CMAKE_ARGV4)
    set(OUTPUT_DIR ${CMAKE_ARGV4})
else()
    set(OUTPUT_DIR "build")
endif()

# 定义平台映射
if(${PLATFORM_ARG} STREQUAL "linux")
    set(GOOS "linux")
    set(GOARCH "amd64")
    set(OUTPUT_NAME "azhot_linux_amd64")
elseif(${PLATFORM_ARG} STREQUAL "windows")
    set(GOOS "windows")
    set(GOARCH "amd64")
    set(OUTPUT_NAME "azhot_windows_amd64.exe")
elseif(${PLATFORM_ARG} STREQUAL "darwin" OR ${PLATFORM_ARG} STREQUAL "macos")
    set(GOOS "darwin")
    set(GOARCH "amd64")
    set(OUTPUT_NAME "azhot_darwin_amd64")
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
file(MAKE_DIRECTORY ${OUTPUT_DIR})

# 执行交叉编译
find_program(GO_CMD go)
if(NOT GO_CMD)
    message(FATAL_ERROR "go command not found. Please install Go.")
endif()

# 设置环境变量并执行构建
set(ENV_COMMAND)
if(WIN32)
    set(ENV_COMMAND cmd /c "set GOOS=${GOOS} && set GOARCH=${GOARCH} && set CGO_ENABLED=0 && go build -ldflags=\"-s -w\" -o ${OUTPUT_DIR}/${OUTPUT_NAME} .")
else()
    set(ENV_COMMAND env GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=0 go build -ldflags="-s -w" -o ${OUTPUT_DIR}/${OUTPUT_NAME} .)
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

message(STATUS "Successfully built for ${PLATFORM_ARG} -> ${OUTPUT_DIR}/${OUTPUT_NAME}")

# 可选：如果指定了package参数，创建zip包
if(CMAKE_ARGV5 AND ${CMAKE_ARGV5} STREQUAL "package")
    # 创建临时目录用于打包
    set(PACKAGE_TEMP_DIR "${OUTPUT_DIR}/temp_package_${GOOS}_${GOARCH}")
    file(MAKE_DIRECTORY ${PACKAGE_TEMP_DIR})
    
    # 复制二进制文件
    file(COPY ${OUTPUT_DIR}/${OUTPUT_NAME} DESTINATION ${PACKAGE_TEMP_DIR})
    
    # 复制必需的配置文件
    if(EXISTS "${CMAKE_CURRENT_SOURCE_DIR}/.env.example")
        execute_process(COMMAND ${CMAKE_COMMAND} -E copy_if_different 
            "${CMAKE_CURRENT_SOURCE_DIR}/.env.example" 
            "${PACKAGE_TEMP_DIR}/.env")
    endif()
    
    if(EXISTS "${CMAKE_CURRENT_SOURCE_DIR}/README.md")
        execute_process(COMMAND ${CMAKE_COMMAND} -E copy_if_different 
            "${CMAKE_CURRENT_SOURCE_DIR}/README.md" 
            "${PACKAGE_TEMP_DIR}/")
    endif()
    
    if(EXISTS "${CMAKE_CURRENT_SOURCE_DIR}/LICENSE")
        execute_process(COMMAND ${CMAKE_COMMAND} -E copy_if_different 
            "${CMAKE_CURRENT_SOURCE_DIR}/LICENSE" 
            "${PACKAGE_TEMP_DIR}/")
    endif()
    
    # 创建zip包
    set(ZIP_NAME "azhot_${GOOS}_${GOARCH}.zip")
    execute_process(
        COMMAND ${CMAKE_COMMAND} -E chdir ${PACKAGE_TEMP_DIR} 
        ${CMAKE_COMMAND} -E tar "cf" "../${ZIP_NAME}" "--format=zip" "."
        WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
        RESULT_VARIABLE ZIP_RESULT
    )
    
    if(NOT ZIP_RESULT EQUAL 0)
        message(WARNING "Failed to create zip package for ${PLATFORM_ARG}")
    else()
        message(STATUS "Successfully created package: ${OUTPUT_DIR}/${ZIP_NAME}")
    endif()
    
    # 清理临时目录
    file(REMOVE_RECURSE ${PACKAGE_TEMP_DIR})
endif()