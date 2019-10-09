package dude

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/mattn/go-zglob"
	"log"
)

func RunFmt() error {

	hclFiles, err := zglob.Glob("**/*.hcl")
	if err != nil {
		return err
	}

	parser := hclparse.NewParser()

	for _, v := range hclFiles {
		hclParseTree, err := parser.ParseHCLFile(v)
		if err != nil {
			panic(err)
		}
		log.Println(spew.Sdump(hclParseTree))
	}

	return nil

}
