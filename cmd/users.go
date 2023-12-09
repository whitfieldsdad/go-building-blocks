package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/whitfieldsdad/go-building-blocks/pkg/bb"
)

var usersCmd = &cobra.Command{
	Use:  "users",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		users, err := bb.ListUsers()
		if err != nil {
			log.Fatalf("Failed to list users: %v", err)
		}
		for _, user := range users {
			printJSON(user)
		}
	},
}

func init() {
	rootCmd.AddCommand(usersCmd)
}
