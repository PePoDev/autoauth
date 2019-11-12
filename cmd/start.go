package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start autoauth",
	Long: fmt.Sprint("AutoAuth is a CLI to set automatic authentication for Internet Login Portal\n",
		"Documents can be found here https://www.github.com/PePoDev/autoauth"),
	SuggestFor: []string{"run"},
	Example:    "autoauth start --deteach --save --service kmitl --user 59050220 --password s3cr3t_p@ssw0rd --interval 5",
	Run: func(cmd *cobra.Command, args []string) {
		if false {
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
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringP("name", "n", "", "name of preset you want to start")
	startCmd.Flags().BoolP("detached", "d", false, "run AutoAuth in the background")
	rootCmd.MarkFlagRequired("region")
}
