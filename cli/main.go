package main

import (
	// "bufio"
	// "fmt"
	// "os"
	// "strings"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "cli",
		Short: "CLI for interacting with the API",
		Long:  `CLI for interacting with the API. Supports login, refresh, logout, and user management.`,
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
