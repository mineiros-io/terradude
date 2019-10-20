package config

import (
	"github.com/rs/zerolog/log"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func DecodeBackendBlock(hclconfigs []*Config, ctx *hcl.EvalContext) (*hclwrite.Block, hcl.Diagnostics) {
	var diags   hcl.Diagnostics
	var backend *Backend

	for _, hclconfig := range hclconfigs {
		if hclconfig.Backend != nil {
			backend = hclconfig.Backend
			break
		}
	}

	body := backend.Body
	block := gohcl.EncodeAsBlock(backend, "backend")
	attrs, diags := body.JustAttributes()
	if diags.HasErrors() {
		log.Fatal().Msg(diags.Error())
	}
	for _,attr := range attrs {
		val, err := attr.Expr.Value(ctx)
		if err != nil {
			return nil, err
		}
		block.Body().SetAttributeValue(attr.Name, val)
	}

	return block, diags
}
