package config

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/rs/zerolog/log"
	"sort"
)

func DecodeTerraformBlock(configs []*Config, ctx *hcl.EvalContext) ([]*hclwrite.Block, hcl.Diagnostics) {
	var diags hcl.Diagnostics
	var terraform *Terraform
	var blocks []*hclwrite.Block

	if configs[0] == nil || configs[0].Terraform == nil {
		log.Fatal().
			Msg("terraform block not defined in leaf")
	}

	terraform = configs[0].Terraform

	block := gohcl.EncodeAsBlock(terraform.Module, "module")
	attrs, diags := terraform.Module.Body.JustAttributes()
	if diags.HasErrors() {
		log.Fatal().
			Msg(diags.Error())
	}

	var keys []string
	for a := range attrs {
		keys = append(keys, a)
	}
	sort.Strings(keys)
	for _, k := range keys {
		e := attrs[k].Expr.(hclsyntax.Expression)
		val, err := e.Value(ctx)
		if err != nil {
			t, isScopeTraversalExpr := e.(*hclsyntax.ScopeTraversalExpr)
			if isScopeTraversalExpr {
				block.Body().SetAttributeTraversal(attrs[k].Name, t.Traversal)
			} else {
				log.Error().
					Msg("attrs[k].Expr")
				for _, i := range err {
					log.Error().
						Msg(i.Error())
				}
				return nil, err
			}
		} else {
			block.Body().SetAttributeValue(attrs[k].Name, val)
		}
	}
	blocks = append(blocks, block)

	return blocks, diags
}
