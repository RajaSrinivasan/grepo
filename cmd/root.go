package cmd

import (
	"fmt"
	"os"

	"github.com/RajaSrinivasan/grepo/impl/repo"
	"github.com/spf13/cobra"
)

var cfgFile string
var verbose bool

var rootCmd = &cobra.Command{
	Use:   "grepo",
	Short: "Group of Repositories support",
	Long: `
	grepo supports a project that comprises different repositories.
	`,
	Version: "v0.0.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "Project.yaml", "config file.")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "be verbose")
}

func initConfig() {
	repo.Verbose = verbose
	repo.LoadConfig(cfgFile)
}
