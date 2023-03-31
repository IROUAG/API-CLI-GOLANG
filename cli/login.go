package main

import (
	"github.com/spf13/cobra"
)

func makeLoginCmd() *cobra.Command {
	var email, password string

	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "Login and receive an authentication JWT and refresh token",
		Run: func(cmd *cobra.Command, args []string) {
			// Implement the logic to login and receive JWT and refresh tokens
		},
	}

	loginCmd.Flags().StringVarP(&email, "email", "e", "", "User email")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "User password")

	loginCmd.MarkFlagRequired("email")
	loginCmd.MarkFlagRequired("password")

	return loginCmd
}
