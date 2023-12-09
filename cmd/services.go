package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/whitfieldsdad/go-building-blocks/pkg/bb"
)

var servicesCmd = &cobra.Command{
	Use:  "services",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		opts := getProcessOptions(cmd.Flags())
		svcs, err := bb.ListServices(opts)
		if err != nil {
			log.Fatalf("Failed to list services: %v", err)
		}
		for _, svc := range svcs {
			printJSON(svc)
		}
	},
}

func init() {
	rootCmd.AddCommand(servicesCmd)

	servicesCmd.PersistentFlags().Bool("include-all", false, "Include file")
	servicesCmd.PersistentFlags().Bool("include-file-metadata", false, "Include file metadata")
	servicesCmd.PersistentFlags().Bool("include-file-hashes", false, "Include file hashes")
}
