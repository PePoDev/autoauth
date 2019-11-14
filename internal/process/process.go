package process

import (
	"errors"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/mitchellh/go-ps"
	"github.com/pepodev/xlog"
)

// FindProcess ...
func FindProcess(key string) (pid int, pname string, err error) {
	err = errors.New("process not found")
	ps, _ := ps.Processes()

	for i := range ps {
		if strings.Contains(ps[i].Executable(), key) && ps[i].Pid() != os.Getpid() {
			pid = ps[i].Pid()
			pname = ps[i].Executable()
			err = nil
			xlog.Debugf("pid: %v name: %v err: %v", pid, pname, err)
		}
	}

	return
}

// KillProcessID ...
func KillProcessID(pid int) error {
	kill := exec.Command("taskkill", "/T", "/F", "/PID", strconv.Itoa(pid))
	return kill.Start()
}
