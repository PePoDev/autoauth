package process

import (
	"os"
	"strings"

	"github.com/mitchellh/go-ps"
	"github.com/pepodev/xlog"
)

// FindProcess ...
func FindProcess(key string) (foundPS []ps.Process) {
	ps, _ := ps.Processes()

	for i := range ps {
		if strings.Contains(ps[i].Executable(), key) && ps[i].Pid() != os.Getpid() {
			foundPS = append(foundPS, ps[i])
			xlog.Debugf("pid: %v name: %v", ps[i].Pid(), ps[i].Executable())
		}
	}

	return
}

// KillProcessID ...
func KillProcessID(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		xlog.Fatalf("Error when find process [PID: %v] [msg: %v]", pid, err)
	}
	err = process.Kill()
	return err
}
