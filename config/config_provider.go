package config

import (
	"github.com/rs/zerolog/log"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func DecodeProviderBlocks(configs []*Config, ctx *hcl.EvalContext) ([]*hclwrite.Block, hcl.Diagnostics) {
	var diags hcl.Diagnostics
	// var providers []*Provider

	providers := map[string]*Provider{}
	blocks := []*hclwrite.Block{}

	for _, config := range configs {
		for _, provider := range config.Provider {
			id := provider.Name
			if provider.Alias != nil {
				id += "." + *provider.Alias
			}
			if providers[id] != nil {
				log.Warn().Msgf("ignoring provider %s (redefined)", id)
				continue
			}
			log.Debug().Msgf("provider: %s", id)
			providers[id] = provider
		}
	}

	for _, provider := range providers {
		body := provider.Body
		block := gohcl.EncodeAsBlock(provider, "provider")
		attrs, diags := body.JustAttributes()
		if diags.HasErrors() {
			log.Fatal().Msg(diags.Error())
		}
		for _, attr := range attrs {
			val, err := attr.Expr.Value(ctx)
			if err != nil {
				return nil, err
			}
			block.Body().SetAttributeValue(attr.Name, val)
		}
    blocks = append(blocks, block)
	}

	return blocks, diags
}
