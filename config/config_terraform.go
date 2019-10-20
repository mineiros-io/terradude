package config

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/rs/zerolog/log"
	"sort"
)

func DecodeTerraformBlock(configs []*Config, ctx *hcl.EvalContext) (*hclwrite.Block, hcl.Diagnostics) {
	var diags hcl.Diagnostics
	var terraform *Terraform

	if configs[0] == nil || configs[0].Terraform == nil {
		log.Fatal().Msg("terraform block not defined in leaf")
	}

	terraform = configs[0].Terraform

	block := gohcl.EncodeAsBlock(terraform.Module, "module")
	attrs, diags := terraform.Module.Body.JustAttributes()
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

	return block, diags
}
