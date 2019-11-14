package cmd

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"github.com/pepodev/xlog"

	"github.com/pepodev/autoauth/internal/message"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create preset file for autoauth configuration",
	Long:    message.GetWelcome(),
	Example: "autoauth create",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter Username: ")
		username, _ := reader.ReadString('\n')

		fmt.Print("Enter Password: ")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		if err == nil {
			fmt.Println("\nPassword typed: " + string(bytePassword))
		}
		password := string(bytePassword)

		xlog.Infof("%s %s", username, password)
	},
}

func init() {

}
