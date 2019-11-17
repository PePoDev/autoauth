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
	Example: "autoauth stop -p [pid]",
	Run: func(cmd *cobra.Command, args []string) {
		fileName, _ := cmd.Flags().GetString("file")
		key, _ := cmd.Flags().GetString("key")
		pid, _ := cmd.Flags().GetInt("process")
		all, _ := cmd.Flags().GetBool("all")
		viper.SetDefault("key", key)

		ps := process.FindProcess("autoauth")
		if len(ps) == 0 {
			xlog.Fatalf("%v", "process not found")
		}

		if all {
			for _, p := range ps {
				if err := process.KillProcessID(p.Pid()); err != nil {
					xlog.Fatalf("%v", err)
				}
			}
			xlog.Info("stopped all autoauth process successful")
		} else {
			if pid == 0 {
				xlog.Fatalf("%v", "need flag `--process` or `-p` that you want to stop")
			}
			for _, p := range ps {
				if p.Pid() == pid {
					if err := process.KillProcessID(pid); err != nil {
						xlog.Fatalf("%v", err)
					}
					xlog.Infof("stopped autoauth process id %v successful", pid)
					return
				}
			}
		}

		preset := auth.LoadPresetFromPath(utils.GetWorkingDirectory(), fileName)
		if err := preset.RequestLogout(); err != nil {
			xlog.Fatalf("%v", err)
		}

	},
}

func init() {
	stopCmd.Flags().StringP("file", "f", "autoauth.yml", "file of preset you want to start auto auth")
	stopCmd.Flags().StringP("key", "k", "", "key to decrypt your data from config file")
	stopCmd.Flags().IntP("process", "p", 0, "process id you want to stop")
	stopCmd.Flags().BoolP("all", "a", false, "stop all processes")
}
