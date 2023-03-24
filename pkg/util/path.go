package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindProjectRootPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// 从当前目录开始，向上查找
	for {
		if dir == "/" {
			return "", fmt.Errorf("not found project root path")
		}

		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		dir = filepath.Dir(dir)
	}
}
