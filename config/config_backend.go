package config

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/rs/zerolog/log"
	"sort"
)

func DecodeBackendBlock(hclconfigs []*Config, ctx *hcl.EvalContext) (*hclwrite.Block, hcl.Diagnostics) {
	var diags hcl.Diagnostics
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
	var keys []string
	for a := range attrs {
		keys = append(keys, a)
	}
	sort.Strings(keys)
	for _, k := range keys {
		val, err := attrs[k].Expr.Value(ctx)
		if err != nil {
			return nil, err
		}
		block.Body().SetAttributeValue(attrs[k].Name, val)
	}
	tfblock := hclwrite.NewBlock("terraform", nil)
	tfblock.Body().AppendBlock(block)

	return tfblock, diags
}
