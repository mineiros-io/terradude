package cli

import (
	"github.com/mineiros-io/terradude/config"
	"github.com/mineiros-io/terradude/dude"
	"github.com/mineiros-io/terradude/util"
	"github.com/rs/zerolog/log"
	"os"
)

func Run() {
	leafs, _ := util.FindLeafFiles(config.DefaultConfigFileBaseName, os.Args[1:], nil)

	for _, leaf := range leafs {
		log.Debug().Msgf("found leaf in %s", leaf)
	}

	for _, leaf := range leafs {
		dude.RunFmt(leaf)
	}
}
