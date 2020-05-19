package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Tag each repo",
	Long: `
	foreach repo - git tag
	`,
	Run: Tag,
}

func init() {
	rootCmd.AddCommand(tagCmd)
}

func Tag(cmd *cobra.Command, args []string) {
	if verbose {
		log.Printf("Assign a tag to the repos")
	}
}
