package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "terradude",
	Short: "A terraform wrapper to work with multiple terraform modules",
	Long: `terradude - Terradude is a thin wrapper for Terraform that provides extra tools for working with multiple
Terraform modules, remote state, and locking. For documentation, see https://github.com/mineiros-io/terradude/.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
