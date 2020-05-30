package utils

import (
	"os"

	"github.com/pepodev/xlog"
)

// GetWorkingDirectory will return path of current working directory
func GetWorkingDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		xlog.Fatalf("%v", err)
	}
	xlog.Debugf("current directory: %s", dir)
	return dir
}
