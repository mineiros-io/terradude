package config

import (
	"github.com/hashicorp/hcl/v2"
)

func DecodeTerraformBlock(hclconfigs []*Config) (*Terraform, hcl.Diagnostics) {
	var diags   hcl.Diagnostics
	return hclconfigs[0].Terraform, diags
}
