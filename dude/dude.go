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
		log.Printf(" - %s\n", f)

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
		if config.Terradude != nil {
			log.Printf("     Terradude    = %#v", config.Terradude)
		}
		if config.Terraform != nil {
			log.Printf("     Terraform    = %#v", config.Terraform)
		}
		if config.Backend != nil {
			log.Printf("     Backend      = %#v", config.Backend)
		}
		for _, provider := range config.Provider {
			log.Printf("     Provider     = %#v", provider)
		}
		for _, dependency := range config.Dependency {
			log.Printf("     Dependency   = %#v", dependency)
		}
		if config.Globals != nil {
			g,_ := config.Globals.Body.JustAttributes()
			for _, attr := range g {
				log.Printf("     Globals.%s = %#v", attr.Name, &attr.Expr)
			}
		}
	}

	return nil
}
