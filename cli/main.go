package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "your-cli",
		Short: "CLI for interacting with the API",
		Long:  `CLI for interacting with the API. Supports login, refresh, logout, and user management.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Read user input in a loop
			reader := bufio.NewReader(os.Stdin)
			for {
				fmt.Print("> ")

				// Read the user's input
				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("Error reading input:", err)
					continue
				}

				// Remove newline characters
				input = strings.TrimSuffix(input, "\n")
				input = strings.TrimSuffix(input, "\r")

				// Split input into arguments
				args := strings.Fields(input)

				// Execute the command with the provided arguments
				cmd.SetArgs(args)
				if err := cmd.Execute(); err != nil {
					fmt.Println("Error executing command:", err)
				}
			}
		},
	}

	loginCmd := createLoginCmd()
	refreshCmd := createRefreshCmd()
	logoutCmd := createLogoutCmd()
	usersCmd := createUsersCmd()

	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(refreshCmd)
	rootCmd.AddCommand(logoutCmd)
	rootCmd.AddCommand(usersCmd)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
