package util

import (
	"os"
	"runtime"
)

func OSNewLine() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}

func FileExists(filePath string) bool {
	_, err := os.Lstat(filePath)
	return !os.IsNotExist(err)
}
