package cmd

import (
	"github.com/pepodev/autoauth/internal/message"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:     "stop",
	Short:   "Stop autoauth",
	Long:    message.GetWelcome(),
	Example: "autoauth stop [name]",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
}
