package main

import (
	"os"
	"time"

	"github.com/mineiros-io/terradude/config"
	"github.com/mineiros-io/terradude/dude"
	"github.com/mineiros-io/terradude/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/alecthomas/kingpin.v2"
)

// This variable is set at build time using -ldflags parameters. For more info, see:
// http://stackoverflow.com/a/11355611/483528
//var VERSION string

var (
	app       = kingpin.New("terradude", "A thin wrapper for terraform.")
	debug     = app.Flag("debug", "Enable debug mode.").Bool()
	jsonlog   = app.Flag("jsonlog", "Enable JSON logging.").Bool()
	directory = app.Arg("directory", "Directory to run in.").Default(".").String()
)

func main() {
	kingpin.Version("0.0.1")
	//app.Parse(os.Args[1:])

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if *jsonlog {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	} else {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	}

	leafs, _ := util.FindLeafFiles(config.DefaultConfigFileBaseName, []string{*directory}, nil)

	for _, leaf := range leafs {
		log.Debug().
			Str("file", leaf).
			Msg("found leaf config")
	}

	for _, leaf := range leafs {
		err := dude.RunFmt(leaf)
		if err != nil {
			log.Debug().
				Str("error", err.Error()).
				Msg("found leaf config")
		}
	}
}
