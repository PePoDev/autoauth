package cmd

import (
	"github.com/sirupsen/logrus"

	"github.com/pepodev/autoauth/internal/message"
	"github.com/pepodev/xlog"
	"github.com/spf13/cobra"
)

var author string = "PePoDev <pepo.dev@outlook.com>"
var license string = "MIT"
var version string = "v0.1.0"

// RootCmd is the root of command
var RootCmd = &cobra.Command{
	Use:     "autoauth",
	Long:    message.GetWelcome(),
	Version: version,
}

func init() {
	xlog.DefaultXlogFormatter()
	xlog.SetLevel(logrus.DebugLevel)

	RootCmd.AddCommand(startCmd)
	RootCmd.AddCommand(stopCmd)
	RootCmd.AddCommand(createCmd)
}
