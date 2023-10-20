package utils

import (
	"os"
	"os/exec"

	"github.com/pepodev/xlog"
)

// StartInDetachMode will run application in background mode
func StartInDetachMode() {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		if args[i] == "-d" {
			args[i] = ""
			break
		}
	}
	cmd := exec.Command(os.Args[0], args...)
	err := cmd.Start()
	if err != nil {
		xlog.Fatalf("%v", err)
	}
	xlog.Infof("Process has started with ID %v and still running in detached mode", cmd.Process.Pid)
	xlog.Infof("To stop this process use command `autoauth stop -p %v` to stop this process", cmd.Process.Pid)
	xlog.Info("Or use `autoauth ps` to see all process is running")
	os.Exit(0)
}
