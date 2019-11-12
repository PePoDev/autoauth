package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop a previous process",
	Long: fmt.Sprint("Stop a previous process\n",
		"Documents can be found here https://www.github.com/PePoDev/autoauth"),
	Example: "autoauth stop [name]",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
