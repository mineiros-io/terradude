package cmd

import (
	"github.com/mineiros-io/terradude/config"
	"github.com/mineiros-io/terradude/util"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all found terraform modules",
	Long:  `Searches the given path backwards recursively for terradude.hcl files`,
	Run: func(cmd *cobra.Command, args []string) {
		leafs, _ := util.FindLeafFiles(config.DefaultConfigFileBaseName, args[0:], nil)

		for _, leaf := range leafs {
			log.Debug().Msgf("found leaf in %s", leaf)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
