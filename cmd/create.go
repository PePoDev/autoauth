package cmd

import (
	"github.com/pepodev/autoauth/internal/message"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create preset file for autoauth configuration",
	Long:    message.GetWelcome(),
	Example: "autoauth stop [name]",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
}
