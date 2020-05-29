package cmd

import (
	"log"

	"github.com/RajaSrinivasan/grepo/impl"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Project Status",
	Long: `
	git status of each repo
	`,
	Run: Status,
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func Status(cmd *cobra.Command, args []string) {
	if verbose {
		log.Printf("Status of the individual projectlets")
	}
	impl.Status(repoconfig)
}
