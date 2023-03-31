package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "app",
		Short: "A CLI application to manage users, roles, and groups",
		Long:  `A CLI application to manage users, roles, and groups through a REST API.`,
	}

	loginCmd := makeLoginCmd()
	refreshCmd := makeRefreshCmd()
	logoutCmd := makeLogoutCmd()
	userCmd := makeUserCmd()
	roleCmd := makeRoleCmd()
	groupCmd := makeGroupCmd()

	rootCmd.AddCommand(loginCmd, refreshCmd, logoutCmd, userCmd, roleCmd, groupCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
