package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/whitfieldsdad/go-building-blocks/pkg/bb"
)

var hostCmd = &cobra.Command{
	Use:  "host",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		host, err := bb.GetHost()
		if err != nil {
			log.Fatalf("Failed to get host: %v", err)
		}
		printJSON(host)
	},
}

func init() {
	rootCmd.AddCommand(hostCmd)
}
