package config

import (
	"github.com/zclconf/go-cty/cty"
	"github.com/hashicorp/hcl/v2"
)

type Global struct {
	Name string
	Value cty.Value
}

func DecodeGlobalValues(hclconfigs []*Config) ([]*Global, hcl.Diagnostics) {
	var globals []*Global
	var diags   hcl.Diagnostics

  for _, hclconfig := range hclconfigs {
		if hclconfig.Globals != nil {
			attrs, diag := hclconfig.Globals.Body.JustAttributes()
			if diag.HasErrors() {
				return nil, diags
			}
			diags = append(diags, diag...)
			attrloop: for _, attr := range attrs {
				for _, global := range globals {
					if (global.Name == attr.Name) {
						continue attrloop
					}
				}
				value, diag := attr.Expr.Value(nil)
				diags = append(diags, diag...)
				if diag.HasErrors() {
					return nil, diags
				}
				globals = append(globals, &Global{
					Name: attr.Name,
					Value: value,
				})
			}
		}
	}
	return globals, diags
}
