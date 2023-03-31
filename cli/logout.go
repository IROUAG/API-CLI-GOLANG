package main

import (
	"github.com/spf13/cobra"
)

func makeLogoutCmd() *cobra.Command {
	var accessToken, refreshToken string

	logoutCmd := &cobra.Command{
		Use:   "logout",
		Short: "Delete an authentication JWT and refresh token",
		Run: func(cmd *cobra.Command, args []string) {
			// Implement the logic to delete JWT and refresh tokens
		},
	}

	logoutCmd.Flags().StringVarP(&accessToken, "access_token", "a", "", "Access token to delete")
	logoutCmd.Flags().StringVarP(&refreshToken, "refresh_token", "r", "", "Refresh token to delete")

	logoutCmd.MarkFlagRequired("access_token")
	logoutCmd.MarkFlagRequired("refresh_token")

	return logoutCmd
}
