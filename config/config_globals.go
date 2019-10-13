package config

import (
	"github.com/zclconf/go-cty/cty/gocty"
	"github.com/zclconf/go-cty/cty"
	"github.com/hashicorp/hcl/v2"
)


type Global struct {
	Name string
	Value cty.Value
}

func DecodeGlobalCty(hclconfigs []*Config) (*cty.Value, hcl.Diagnostics) {
	var diags   hcl.Diagnostics
	globals := map[string]cty.Value{}

	for _, hclconfig := range hclconfigs {
		if hclconfig.Globals != nil {
			attrs, diag := hclconfig.Globals.Body.JustAttributes()
			if diag.HasErrors() {
				return nil, diags
			}
			diags = append(diags, diag...)
			for _, attr := range attrs {
				if _, ok := globals[attr.Name] ; ok {
					continue
				}
				value, diag := attr.Expr.Value(nil)
				diags = append(diags, diag...)
				if diag.HasErrors() {
					return nil, diags
				}
				globals[attr.Name] = value
			}
		}
	}

	ctyTypes := map[string]cty.Type{}
	for key, value := range globals {
		ctyTypes[key] = value.Type()
	}
	ctyObject := cty.Object(ctyTypes)
	ctyGlobals, err := gocty.ToCtyValue(globals, ctyObject)
	if err != nil {
		return nil, err.(hcl.Diagnostics)
	}

	return &ctyGlobals, nil
}
