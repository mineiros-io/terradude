package main

import (
	"github.com/mineiros-io/terradude/config"
	"github.com/mineiros-io/terradude/dude"
	"github.com/mineiros-io/terradude/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

// This variable is set at build time using -ldflags parameters. For more info, see:
// http://stackoverflow.com/a/11355611/483528
var VERSION string

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})

	leafs, _ := util.FindLeafFiles(config.DefaultConfigFileBaseName, os.Args[1:], nil)

	for _, leaf := range leafs {
		log.Debug().Msgf("found leaf in %s", leaf)
	}

	for _, leaf := range leafs {
		dude.RunFmt(leaf)
	}
}
