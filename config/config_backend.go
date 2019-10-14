package config

import (
	"github.com/zclconf/go-cty/cty"
	"github.com/hashicorp/hcl/v2"
)

func DecodeBackendBlock(hclconfigs []*Config, globals *cty.Value) (*Backend, hcl.Diagnostics) {
	var diags hcl.Diagnostics

	for _, hclconfig := range hclconfigs {
		if hclconfig.Backend != nil {
			return hclconfig.Backend, diags
		}
	}

	return nil, diags
}
