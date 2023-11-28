package utils

import (
	"path"
	"runtime"
)

func GetCurrentDir() string {
	_, file, _, _ := runtime.Caller(1)
	return path.Dir(file)
}
