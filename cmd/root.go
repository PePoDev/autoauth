package cmd

import (
	"github.com/sirupsen/logrus"

	"github.com/pepodev/autoauth/internal/message"
	"github.com/pepodev/xlog"
	"github.com/spf13/cobra"
)

var version string = "v1.2.1"

// RootCmd is the root of command
var RootCmd = &cobra.Command{
	Use:     "autoauth",
	Long:    message.GetWelcome(),
	Version: version,
}

func init() {
	xlog.DefaultXlogFormatter()
	xlog.SetLevel(logrus.InfoLevel)

	RootCmd.AddCommand(startCmd)
	RootCmd.AddCommand(stopCmd)
	RootCmd.AddCommand(createCmd)
	RootCmd.AddCommand(psCmd)
}
