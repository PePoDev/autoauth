package cmd

import (
	"github.com/pepodev/autoauth/internal/message"
	"github.com/spf13/cobra"
)

var psCmd = &cobra.Command{
	Use:     "ps",
	Short:   "Show all process of autoauth",
	Long:    message.GetWelcome(),
	Example: "autoauth ps",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {

}
