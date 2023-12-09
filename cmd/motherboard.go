package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/whitfieldsdad/go-building-blocks/pkg/bb"
)

var motherboardCmd = &cobra.Command{
	Use:  "motherboard",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		mobo, err := bb.GetMotherboard()
		if err != nil {
			log.Fatalf("Failed to get motherboard: %v", err)
		}
		printJSON(mobo)
	},
}

func init() {
	rootCmd.AddCommand(motherboardCmd)
}
