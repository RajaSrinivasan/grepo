package cmd

import (
	"log"

	"github.com/RajaSrinivasan/grepo/impl"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Diff from where we started",
	Long: `
	foreach repo - provide a difference from the original pull.

	Since the public repo's are considered not under development, this command applies only to the private project group.

	`,
	Run: Diff,
}

func init() {
	rootCmd.AddCommand(diffCmd)
}

func Diff(cmd *cobra.Command, args []string) {
	if verbose {
		log.Printf("Show the difference")
	}
	impl.Verbose = verbose
	impl.Diff(repoconfig)
}
