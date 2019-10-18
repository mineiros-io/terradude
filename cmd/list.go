package cmd

import (
	"github.com/mineiros-io/terradude/config"
	"github.com/mineiros-io/terradude/dude"
	"github.com/mineiros-io/terradude/util"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all found terraform modules",
	Long:  `Searches the given path backwards recursively for terradude.hcl files`,
	Run: func(cmd *cobra.Command, args []string) {

		leafs, _ := util.FindLeafFiles(config.DefaultConfigFileBaseName, os.Args[1:], nil)

		for _, leaf := range leafs {
			log.Debug().Msgf("found leaf in %s", leaf)
		}

		for _, leaf := range leafs {
			dude.RunFmt(leaf)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
