package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Project Status",
	Long: `
	Status of the project followed by git status of each repo
	`,
	Run: Status,
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func Status(cmd *cobra.Command, args []string) {
	if verbose {
		log.Printf("Status of the overall project and the individual projectlets")
	}
}
