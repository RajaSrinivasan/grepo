package cmd

import (
	"log"

	"github.com/RajaSrinivasan/grepo/impl"
	"github.com/RajaSrinivasan/grepo/impl/repo"
	"github.com/spf13/cobra"
)

var force bool
var build bool
var branch string

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
	initCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "Force init. If previously initialized - wipeout.")
	initCmd.PersistentFlags().BoolVarP(&build, "build", "b", false, "Perform initial build.")
	initCmd.PersistentFlags().StringVarP(&branch, "new-branch", "n", repo.NoneBranch, "Create a new branch upon checkout")
	rootCmd.AddCommand(initCmd)
}

func Init(cmd *cobra.Command, args []string) {
	if verbose {
		log.Printf("Initialize force=%v", force)
	}
	impl.Verbose = verbose
	impl.Init(repoconfig, force, build, branch)
}
