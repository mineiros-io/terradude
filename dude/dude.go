package dude

import (
	"fmt"
	"path/filepath"
	"github.com/zclconf/go-cty/cty"
	"github.com/mineiros-io/terradude/config"
	"github.com/rs/zerolog/log"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func RunFmt(file string) error {
	log.Info().Msgf("> start processing %v", filepath.Dir(file))

	hclconfigs, terradude, diags := config.LoadConfigs(file)
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
	ctx.Variables["terradude"] = *terradude

	log.Info().Msgf("+ creating backend config for %s", filepath.Dir(file))
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
	log.Info().Msgf("< finished processing %v", filepath.Dir(file))
	return nil
}
