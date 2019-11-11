package main

import (
	"AutoAuthen/pkg/auth"
	"AutoAuthen/pkg/auth/options"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var godaemon = flag.Bool("d", false, "run app as a daemon with -d=true")

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}
	if *godaemon {
		args := os.Args[1:]
		for i := 0; i < len(args); i++ {
			if args[i] == "-d=true" {
				args[i] = "-d=false"
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
