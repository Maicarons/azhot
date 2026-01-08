#!/usr/bin/env cmake -P
# 这是一个辅助CMake脚本来构建所有平台并打包
# 使用方法: cmake -P cmake_scripts/build_all_platforms.cmake [output_dir]

cmake_minimum_required(VERSION 3.12)

# 获取输出目录参数（可选），默认为package
if(CMAKE_ARGV3)
    set(OUTPUT_DIR ${CMAKE_ARGV3})
else()
    set(OUTPUT_DIR "package")
endif()

# 定义要构建的平台列表
set(PLATFORMS linux-amd64 windows-amd64 darwin-amd64)
set(LINUX_AMD64 "linux,amd64,azhot_linux_amd64")
set(WINDOWS_AMD64 "windows,amd64,azhot_windows_amd64.exe")
set(DARWIN_AMD64 "darwin,amd64,azhot_darwin_amd64")

# 创建输出目录
file(MAKE_DIRECTORY ${OUTPUT_DIR})

# 执行交叉编译
find_program(GO_CMD go)
if(NOT GO_CMD)
    message(FATAL_ERROR "go command not found. Please install Go.")
endif()

# 函数：构建单个平台
function(build_platform GOOS GOARCH OUTPUT_NAME)
    message(STATUS "Building for ${GOOS}/${GOARCH}...")
    
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
        message(FATAL_ERROR "Build failed for ${GOOS}/${GOARCH}\nOutput: ${BUILD_OUTPUT}\nError: ${BUILD_ERROR}")
    endif()

    message(STATUS "Successfully built for ${GOOS}/${GOARCH} -> ${OUTPUT_DIR}/${OUTPUT_NAME}")
endfunction()

# 构建所有平台
build_platform("linux" "amd64" "azhot_linux_amd64")
build_platform("windows" "amd64" "azhot_windows_amd64.exe")
build_platform("darwin" "amd64" "azhot_darwin_amd64")

# 打包所有平台
message(STATUS "Packaging all platforms...")

# 为Linux创建临时目录并打包
set(LINUX_TEMP_DIR "${OUTPUT_DIR}/temp_linux")
file(MAKE_DIRECTORY ${LINUX_TEMP_DIR})
file(COPY ${OUTPUT_DIR}/azhot_linux_amd64 DESTINATION ${LINUX_TEMP_DIR})
if(EXISTS "${CMAKE_CURRENT_SOURCE_DIR}/.env.example")
    execute_process(COMMAND ${CMAKE_COMMAND} -E copy_if_different 
        "${CMAKE_CURRENT_SOURCE_DIR}/.env.example" 
        "${LINUX_TEMP_DIR}/.env")
endif()
if(EXISTS "${CMAKE_CURRENT_SOURCE_DIR}/README.md")
    execute_process(COMMAND ${CMAKE_COMMAND} -E copy_if_different 
        "${CMAKE_CURRENT_SOURCE_DIR}/README.md" 
        "${LINUX_TEMP_DIR}/")
endif()
if(EXISTS "${CMAKE_CURRENT_SOURCE_DIR}/LICENSE")
    execute_process(COMMAND ${CMAKE_COMMAND} -E copy_if_different 
        "${CMAKE_CURRENT_SOURCE_DIR}/LICENSE" 
        "${LINUX_TEMP_DIR}/")
endif()

# 创建Linux zip包
execute_process(
    COMMAND ${CMAKE_COMMAND} -E chdir ${LINUX_TEMP_DIR} 
    ${CMAKE_COMMAND} -E tar "cf" "${OUTPUT_DIR}/azhot_linux_amd64.zip" "--format=zip" "."
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    RESULT_VARIABLE ZIP_RESULT
)
if(NOT ZIP_RESULT EQUAL 0)
    message(WARNING "Failed to create zip package for Linux")
endif()
file(REMOVE_RECURSE ${LINUX_TEMP_DIR})

# 为Windows创建临时目录并打包
set(WINDOWS_TEMP_DIR "${OUTPUT_DIR}/temp_windows")
file(MAKE_DIRECTORY ${WINDOWS_TEMP_DIR})
file(COPY ${OUTPUT_DIR}/azhot_windows_amd64.exe DESTINATION ${WINDOWS_TEMP_DIR})
if(EXISTS "${CMAKE_CURRENT_SOURCE_DIR}/.env.example")
    execute_process(COMMAND ${CMAKE_COMMAND} -E copy_if_different 
        "${CMAKE_CURRENT_SOURCE_DIR}/.env.example" 
        "${WINDOWS_TEMP_DIR}/.env")
endif()
if(EXISTS "${CMAKE_CURRENT_SOURCE_DIR}/README.md")
    execute_process(COMMAND ${CMAKE_COMMAND} -E copy_if_different 
        "${CMAKE_CURRENT_SOURCE_DIR}/README.md" 
        "${WINDOWS_TEMP_DIR}/")
endif()
if(EXISTS "${CMAKE_CURRENT_SOURCE_DIR}/LICENSE")
    execute_process(COMMAND ${CMAKE_COMMAND} -E copy_if_different 
        "${CMAKE_CURRENT_SOURCE_DIR}/LICENSE" 
        "${WINDOWS_TEMP_DIR}/")
endif()

# 创建Windows zip包
execute_process(
    COMMAND ${CMAKE_COMMAND} -E chdir ${WINDOWS_TEMP_DIR} 
    ${CMAKE_COMMAND} -E tar "cf" "${OUTPUT_DIR}/azhot_windows_amd64.zip" "--format=zip" "."
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    RESULT_VARIABLE ZIP_RESULT
)
if(NOT ZIP_RESULT EQUAL 0)
    message(WARNING "Failed to create zip package for Windows")
endif()
file(REMOVE_RECURSE ${WINDOWS_TEMP_DIR})

# 为macOS创建临时目录并打包
set(DARWIN_TEMP_DIR "${OUTPUT_DIR}/temp_darwin")
file(MAKE_DIRECTORY ${DARWIN_TEMP_DIR})
file(COPY ${OUTPUT_DIR}/azhot_darwin_amd64 DESTINATION ${DARWIN_TEMP_DIR})
if(EXISTS "${CMAKE_CURRENT_SOURCE_DIR}/.env.example")
    execute_process(COMMAND ${CMAKE_COMMAND} -E copy_if_different 
        "${CMAKE_CURRENT_SOURCE_DIR}/.env.example" 
        "${DARWIN_TEMP_DIR}/.env")
endif()
if(EXISTS "${CMAKE_CURRENT_SOURCE_DIR}/README.md")
    execute_process(COMMAND ${CMAKE_COMMAND} -E copy_if_different 
        "${CMAKE_CURRENT_SOURCE_DIR}/README.md" 
        "${DARWIN_TEMP_DIR}/")
endif()
if(EXISTS "${CMAKE_CURRENT_SOURCE_DIR}/LICENSE")
    execute_process(COMMAND ${CMAKE_COMMAND} -E copy_if_different 
        "${CMAKE_CURRENT_SOURCE_DIR}/LICENSE" 
        "${DARWIN_TEMP_DIR}/")
endif()

# 创建macOS zip包
execute_process(
    COMMAND ${CMAKE_COMMAND} -E chdir ${DARWIN_TEMP_DIR} 
    ${CMAKE_COMMAND} -E tar "cf" "${OUTPUT_DIR}/azhot_darwin_amd64.zip" "--format=zip" "."
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    RESULT_VARIABLE ZIP_RESULT
)
if(NOT ZIP_RESULT EQUAL 0)
    message(WARNING "Failed to create zip package for macOS")
endif()
file(REMOVE_RECURSE ${DARWIN_TEMP_DIR})

message(STATUS "All builds and packages completed successfully in ${OUTPUT_DIR}/")