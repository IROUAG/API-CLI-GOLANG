package main

import (
	"github.com/spf13/cobra"
)

func makeRefreshCmd() *cobra.Command {
	var refreshToken string

	refreshCmd := &cobra.Command{
		Use:   "refresh",
		Short: "Refresh an authentication JWT using a refresh token",
		Run: func(cmd *cobra.Command, args []string) {
			// Implement the logic to refresh JWT using the refresh token
		},
	}

	refreshCmd.Flags().StringVarP(&refreshToken, "refresh_token", "r", "", "Refresh token")

	refreshCmd.MarkFlagRequired("refresh_token")

	return refreshCmd
}
