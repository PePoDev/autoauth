package cmd

import (
	"github.com/pepodev/autoauth/internal/message"
	"github.com/pepodev/autoauth/internal/process"
	"github.com/pepodev/xlog"
	"github.com/spf13/cobra"
)

var psCmd = &cobra.Command{
	Use:     "ps",
	Short:   "Show all process of autoauth",
	Long:    message.GetWelcome(),
	Example: "autoauth ps",
	Run: func(cmd *cobra.Command, args []string) {
		ps := process.FindProcess("autoauth")
		xlog.Infof("Found %v items", len(ps))
		for _, p := range ps {
			xlog.Infof("process id %v", p.Pid())
		}
	},
}

func init() {

}
