package dude

import (
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/mineiros-io/terradude/util"
	"github.com/mineiros-io/terradude/config"
	"log"
)

func RunFmt(file string) error {

	files := util.SearchUp(file)

	var config config.Config

	for _, e := range files {

		hclsimple.DecodeFile(e, nil, &config)

		log.Printf("Configuration is %#v", config)
	}

	return nil

}
