package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/pepodev/xlog"
	"github.com/spf13/cobra"
)

var author string = "PePoDev <pepo.dev@outlook.com>"
var license string = "MIT"
var version string = "v0.1.0"

var rootCmd = &cobra.Command{
	Use: "autoauth",
	Long: fmt.Sprint("AutoAuth is a CLI to set automatic authentication for Internet Login Portal\n",
		"Documents can be found here https://www.github.com/PePoDev/autoauth"),
	Version: version,
}

func init() {
	xlog.DefaultXlogFormatter()
	xlog.SetLevel(logrus.InfoLevel)
}

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
