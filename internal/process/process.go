package process

import (
	"os"
	"strings"

	"github.com/mitchellh/go-ps"
	"github.com/pepodev/xlog"
)

// FindProcess search by name and return processes list
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

// KillProcessID kill process by id
func KillProcessID(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		xlog.Fatalf("Error when find process [PID: %v] [msg: %v]", pid, err)
	}
	err = process.Kill()
	return err
}
