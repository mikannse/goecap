package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

const (
	OutputDir = "./pcaps"
)

// GetLastIndex 获取最新的文件名序号
func GetLastIndex() int {
	lastIndexFile := filepath.Join(OutputDir, "last_index.txt")
	var index int
	if data, err := os.ReadFile(lastIndexFile); err == nil {
		index, _ = strconv.Atoi(string(data))
	} else {
		index = 1
	}
	return index
}

// WriteLastIndex 写入最新的文件名序号
func WriteLastIndex(index int) {
	lastIndexFile := filepath.Join(OutputDir, "last_index.txt")
	if err := os.WriteFile(lastIndexFile, []byte(strconv.Itoa(index)), 0644); err != nil {
		fmt.Printf("Failed to write last index: %v\n", err)
	}
}
