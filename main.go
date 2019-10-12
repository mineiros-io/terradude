package main

import (
	"os"
	"log"
	"github.com/mineiros-io/terradude/config"
	"github.com/mineiros-io/terradude/util"
	"github.com/mineiros-io/terradude/dude"
)


// This variable is set at build time using -ldflags parameters. For more info, see:
// http://stackoverflow.com/a/11355611/483528
var VERSION string

func main() {
	leafs, _ := util.FindLeafFiles(config.DefaultConfigPath, os.Args[1:], nil)

	for _, leaf := range leafs {
		log.Printf("start processing %v\n", leaf)
		dude.RunFmt(leaf)
		log.Printf("finished processing %v\n", leaf)
	}
}
