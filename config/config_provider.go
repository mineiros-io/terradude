package config

import (
	"sort"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/rs/zerolog/log"
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
		blocks = append(blocks, block)
	}

	return blocks, diags
}
