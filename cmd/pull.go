package cmd

import (
	"log"

	"github.com/RajaSrinivasan/grepo/impl"
	"github.com/spf13/cobra"
)

var all_groups bool

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull for each repo",
	Long: `
	foreach repo - git pull

	Clone option - clone the repo.
	               Public repos - only the branch/tag and detach HEAD
				   Private repos - clone, checkout. other branches left accessible
	`,
	Run: Pull,
}

func init() {
	rootCmd.AddCommand(pullCmd)
	pullCmd.PersistentFlags().BoolVarP(&all_groups, "all-groups", "a", false, "All project groups - public and private (default)")

}

func Pull(cmd *cobra.Command, args []string) {
	if verbose {
		log.Printf("Pull the repos")
	}
	impl.Verbose = verbose
	impl.Pull(repoconfig, all_groups, build)
}
