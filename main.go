package main

import (
	"AutoAuthen/pkg/auth"
	"AutoAuthen/pkg/auth/options"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var detached = flag.Bool("d", false, "run app as a daemon with -d=true")

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}
	if *detached {
		args := os.Args[1:]
		for i := 0; i < len(args); i++ {
			if args[i] == "-d" {
				args[i] = ""
				break
			}
		}
		cmd := exec.Command(os.Args[0], args...)
		cmd.Start()
		fmt.Println("[PID]", cmd.Process.Pid)
		os.Exit(0)
	}
}

func main() {
	auth.LoginWithOption(options.KMITL{})
}
