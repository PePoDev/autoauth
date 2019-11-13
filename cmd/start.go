package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pepodev/autoauth/pkg/auth"

	"github.com/pepodev/xlog"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start autoauth",
	Long: fmt.Sprint("AutoAuth is a CLI to set automatic authentication for Internet Login Portal\n",
		"Documents can be found here https://www.github.com/PePoDev/autoauth"),
	SuggestFor: []string{"run"},
	Example:    "autoauth start -d -f config.yml -k my_s3cr3t_k3y",
	Run: func(cmd *cobra.Command, args []string) {
		// Get all variable from command
		isDetached, _ := cmd.Flags().GetBool("detach")
		fileName, _ := cmd.Flags().GetString("file")
		key, _ := cmd.Flags().GetString("key")

		viper.SetDefault("key", key)

		// Detech flag to run in detach mode
		if isDetached {
			args := os.Args[1:]
			for i := 0; i < len(args); i++ {
				if args[i] == "-d" {
					args[i] = ""
					break
				}
			}
			cmd := exec.Command(os.Args[0], args...)
			cmd.Start()
			xlog.Infof("[PID] %v", cmd.Process.Pid)
			os.Exit(0)
		}

		// Get current working directory
		dir, err := os.Getwd()
		if err != nil {
			xlog.Fatalf("%v", err)
		}
		xlog.Debugf("current directory: %s", dir)

		// Read config file from default file name in working directory
		viper.SetConfigFile(fileName)
		viper.AddConfigPath(dir)
		if err := viper.ReadInConfig(); err != nil {
			xlog.Fatalf("fatal error config file: %s \n", err)
		}

		xlog.Info("AutoAuth Started")
		auth.StartAutoLogin()
	},
}

func init() {
	startCmd.Flags().BoolP("detach", "d", false, "start autoauth in the detach mode")
	startCmd.Flags().StringP("file", "f", "autoauth.yml", "file of preset you want to start auto auth")
	startCmd.Flags().StringP("key", "k", "", "key to decrypt your data from config file")
}
