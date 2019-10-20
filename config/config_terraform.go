package config

import (
	"github.com/rs/zerolog/log"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
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
	for _,attr := range attrs {
		val, err := attr.Expr.Value(ctx)
		if err != nil {
			return nil, err
		}
		block.Body().SetAttributeValue(attr.Name, val)
	}

	return block, diags
}
