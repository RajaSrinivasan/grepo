package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push for each repo",
	Long: `
	foreach repo - git push
	`,
	Run: Push,
}

func init() {
	rootCmd.AddCommand(pushCmd)
}

func Push(cmd *cobra.Command, args []string) {
	if verbose {
		log.Printf("Push the repos")
	}
}
