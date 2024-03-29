package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/pepodev/autoauth/internal/message"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create preset file for autoauth configuration",
	Long:    message.GetWelcome(),
	Example: "autoauth create",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Name: ")
		username, _ := reader.ReadString('\n')

		fmt.Print("Key: ")
		byteKey, err := term.ReadPassword(int(syscall.Stdin))
		if err == nil {
			fmt.Println("\nKey typed: " + string(byteKey))
		}
		key := string(byteKey)

		log.Printf("%s %s\n", username, key)
	},
}
