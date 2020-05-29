package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var forcePush bool
var message string

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
	pushCmd.PersistentFlags().BoolVarP(&forcePush, "force", "f", false, "Push all repositories - even if no changes identified.")
	pushCmd.PersistentFlags().StringVarP(&message, "message", "m", "", "Commit message to use on all repositories")
}

func Push(cmd *cobra.Command, args []string) {
	if verbose {
		log.Printf("Push the repos")
	}
}
