package dude

import (
	"fmt"
	"github.com/zclconf/go-cty/cty"
	"github.com/mineiros-io/terradude/config"
	"github.com/rs/zerolog/log"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func RunFmt(file string) error {

	hclconfigs, diags := config.LoadConfigs(file)
	if diags.HasErrors() {
		log.Fatal().Msg(diags.Error())
	}

	// terraform, diags := config.DecodeTerraformBlock(hclconfigs)
	// if diags.HasErrors() {
	// 	log.Fatal().Msg(diags.Error())
	// }

	globals, diags := config.DecodeGlobalCty(hclconfigs)
	if diags.HasErrors() {
		log.Fatal().Msg(diags.Error())
	}

	backend, diags := config.DecodeBackendBlock(hclconfigs, globals)
	if diags.HasErrors() {
		log.Fatal().Msg(diags.Error())
	}

	ctx := &hcl.EvalContext{}
	ctx.Variables = map[string]cty.Value{}
	ctx.Variables["global"] = *globals
	ctx.Variables["terradude"] = *globals

  f := hclwrite.NewEmptyFile()
  b := gohcl.EncodeAsBlock(backend, "backend")
	attrs, _ := backend.Body.JustAttributes()
	for _,attr := range attrs {
		val, err := attr.Expr.Value(ctx)
		if err != nil {
			panic(err)
		}
		b.Body().SetAttributeValue(attr.Name, val)
	}

	f.Body().AppendBlock(b)
	fmt.Printf(string(f.Bytes()))
	return nil
}
