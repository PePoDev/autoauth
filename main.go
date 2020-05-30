package main

import (
	"log"

	"github.com/pepodev/autoauth/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalf("%v", err.Error())
	}
}
