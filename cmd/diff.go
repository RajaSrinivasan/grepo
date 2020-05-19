package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Diff from where we started",
	Long: `
	foreach repo - provide a difference from the original pull
	`,
	Run: Push,
}

func init() {
	rootCmd.AddCommand(diffCmd)
}

func Diff(cmd *cobra.Command, args []string) {
	if verbose {
		log.Printf("Show the difference")
	}
}
