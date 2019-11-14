package cmd

import (
	"github.com/pepodev/autoauth/internal/message"
	"github.com/pepodev/autoauth/internal/process"
	"github.com/pepodev/autoauth/internal/utils"
	"github.com/pepodev/autoauth/pkg/auth"
	"github.com/pepodev/xlog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var stopCmd = &cobra.Command{
	Use:     "stop",
	Short:   "Stop autoauth",
	Long:    message.GetWelcome(),
	Example: "autoauth stop [name]",
	Run: func(cmd *cobra.Command, args []string) {
		fileName, _ := cmd.Flags().GetString("file")
		key, _ := cmd.Flags().GetString("key")
		viper.SetDefault("key", key)

		preset := auth.LoadPresetFromPath(utils.GetWorkingDirectory(), fileName)

		pid, _, err := process.FindProcess("autoauth")
		if err != nil {
			xlog.Fatalf("%v", err)
		}

		if err := process.KillProcessID(pid); err != nil {
			xlog.Fatalf("%v", err)
		}

		preset.RequestLogout()
		xlog.Infof("stopped autoauth process id %v successful", pid)
	},
}

func init() {
	stopCmd.Flags().StringP("file", "f", "autoauth.yml", "file of preset you want to start auto auth")
	stopCmd.Flags().StringP("key", "k", "", "key to decrypt your data from config file")
}
