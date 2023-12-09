package cmd

import (
	"github.com/spf13/pflag"
	"github.com/whitfieldsdad/go-building-blocks/pkg/bb"
)

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
