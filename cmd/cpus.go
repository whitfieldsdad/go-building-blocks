package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/whitfieldsdad/go-building-blocks/pkg/bb"
)

var cpusCmd = &cobra.Command{
	Use:  "cpus",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cpus, err := bb.ListCPUs()
		if err != nil {
			log.Fatalf("Failed to list CPUs: %v", err)
		}
		for _, cpu := range cpus {
			printJSON(cpu)
		}
	},
}

func init() {
	rootCmd.AddCommand(cpusCmd)
}
