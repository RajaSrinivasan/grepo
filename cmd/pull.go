package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull for each repo",
	Long: `
	foreach repo - git pull
	`,
	Run: Pull,
}

func init() {
	rootCmd.AddCommand(pullCmd)
}

func Pull(cmd *cobra.Command, args []string) {
	if verbose {
		log.Printf("Pull the repos")
	}
}
