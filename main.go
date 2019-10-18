package main

import (
	"github.com/mineiros-io/terradude/cmd"
	"github.com/mineiros-io/terradude/cli"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})

	cmd.Execute()
}
