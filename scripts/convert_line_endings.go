package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// convertLineEndings 将文件中的CRLF转换为LF
func convertLineEndings(filePath string) error {
	// 读取文件内容
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// 将CRLF转换为LF
	contentStr := string(content)
	contentStr = strings.ReplaceAll(contentStr, "\r\n", "\n")
	contentStr = strings.ReplaceAll(contentStr, "\r", "\n")

	// 写回文件
	return ioutil.WriteFile(filePath, []byte(contentStr), 0644)
}

// shouldConvert 检查是否应该转换此文件
func shouldConvert(filePath string) bool {
	// 排除不需要处理的目录
	if strings.Contains(filePath, "node_modules") ||
		strings.Contains(filePath, ".git") ||
		strings.Contains(filePath, "vendor") {
		return false
	}

	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".go", ".vue", ".ts", ".js", ".json", ".yaml", ".yml", ".md", ".txt", ".html", ".css", ".py", ".sh", ".conf", ".ini", ".toml", ".cmake":
		return true
	}

	// 检查文件名
	baseName := strings.ToLower(filepath.Base(filePath))
	if baseName == "makefile" || strings.Contains(baseName, ".sh") || strings.Contains(baseName, ".conf") {
		return true
	}

	return false
}

func main() {
	rootDir := "." // 当前目录
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过目录和二进制文件
		if info.IsDir() {
			return nil
		}

		// 检查是否应该转换此文件
		if shouldConvert(path) {
			fmt.Printf("Converting: %s\n", path)
			if err := convertLineEndings(path); err != nil {
				fmt.Printf("Error converting %s: %v\n", path, err)
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("All text files have been converted to LF line endings.")
}
