package main

import (
	"os"

	"github.com/pepodev/autoauth/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
