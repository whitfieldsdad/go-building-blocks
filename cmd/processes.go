package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/whitfieldsdad/go-building-blocks/pkg/bb"
)

var processesCmd = &cobra.Command{
	Use:  "processes",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		opts := getProcessOptions(cmd.Flags())
		processes, err := bb.ListProcesses(opts)
		if err != nil {
			log.Fatalf("Failed to list processes: %v", err)
		}
		for _, process := range processes {
			printJSON(process)
		}
	},
}

func getProcessOptions(flags *pflag.FlagSet) *bb.ProcessOptions {
	includeAll, _ := flags.GetBool("include-all")
	includeFileMetadata, _ := flags.GetBool("include-file-metadata")
	includeFileHashes, _ := flags.GetBool("include-file-hashes")

	return &bb.ProcessOptions{
		IncludeAll:          includeAll,
		IncludeFileMetadata: includeFileMetadata,
		IncludeFileHashes:   includeFileHashes,
	}
}

func init() {
	rootCmd.AddCommand(processesCmd)

	processesCmd.PersistentFlags().Bool("include-all", false, "Include file")
	processesCmd.PersistentFlags().Bool("include-file-metadata", false, "Include file metadata")
	processesCmd.PersistentFlags().Bool("include-file-hashes", false, "Include file hashes")
}
