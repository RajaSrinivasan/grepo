package cmd

import (
	"log"

	"github.com/RajaSrinivasan/grepo/impl"
	"github.com/spf13/cobra"
)

var force bool

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize - setup the workspace",
	Long: `
	Create the workspaces and the directory structure.
	Perform the clone of the repos.
	`,
	Run: Init,
}

func init() {
	initCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "force init. Even if previously initialized.")
	rootCmd.AddCommand(initCmd)
}

func Init(cmd *cobra.Command, args []string) {
	if verbose {
		log.Printf("Initialize force=%v", force)
	}
	impl.Verbose = verbose
	impl.Init(repoconfig, force)
}
