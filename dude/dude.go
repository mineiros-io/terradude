package dude

import (
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/mineiros-io/terradude/util"
	"github.com/mineiros-io/terradude/config"
	"log"
)

func RunFmt(file string) error {

	files := util.SearchUp(file)

	for _, f := range files {
		var config config.Config

		err := hclsimple.DecodeFile(f, nil, &config)
		if err != nil {
			panic(err)
		}

		if f == file && config.Terraform == nil {
			log.Fatalf("terraform block missing in leaf hcl file %s\n", f)
			return nil
		}

		if f != file && config.Terraform != nil {
			log.Fatalf("terraform block in non-leaf hcl file defined %s\n", f)
			return nil
		}

		log.Printf("Configuration is %#v", config)
	}

	return nil
}
