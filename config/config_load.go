package config

import (
	"path/filepath"
	"github.com/zclconf/go-cty/cty"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/mineiros-io/terradude/util"
  "github.com/rs/zerolog/log"
)

func LoadConfigs(configFileName string) ([]*Config, *cty.Value, hcl.Diagnostics) {
	var configs []*Config
	var diags   hcl.Diagnostics
	var file    string

	files     := util.SearchUp(configFileName, DefaultConfigFileBaseName)
	backend   := false
	terradude := map[string]cty.Value{}

	for _, file = range files {
		var config Config
		log.Debug().Msgf("  including config %s", file)

		err := hclsimple.DecodeFile(file, nil, &config)
		if err != nil {
			return nil, nil, err.(hcl.Diagnostics)
		}

		if file == configFileName && config.Terraform == nil {
			diags = append(diags, &hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  "terraform block missing",
				Detail:   "Terradude expects the terraform block to (only) exist in leaf hcl files.",
				Subject: &hcl.Range{
					Filename: file,
					Start:    hcl.InitialPos,
					End:      hcl.InitialPos,
				},
			})
			return nil, nil, diags
		}

		if file != configFileName && config.Terraform != nil {
			diags = append(diags, &hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  "unexpected terraform block found",
				Detail:   "Terradude expects the terraform block to (only) exist in leaf hcl files.",
				Subject: &hcl.Range{
					Filename: file,
					Start:    hcl.InitialPos,
					End:      hcl.InitialPos,
				},
			})
			return nil, nil, diags
		}

		configs = append(configs, &config)

		if config.Backend != nil {
			backend = true
		}

		if config.Terradude != nil && config.Terradude.Version != "" {
			dir    := filepath.Dir(file)
			rel, _ := filepath.Rel(dir, configFileName)

			terradude["module_path"] = cty.StringVal(filepath.Dir(rel))

			log.Debug().Msgf("    setting terradude.module_path = %#v", terradude["module_path"])
			log.Debug().Msgf("    found terradude.version in %s - stop including config files", file)
			break
		}
	}

	if !backend {
		diags = append(diags, &hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "backend block missing",
			Detail:   "Terradude expects a backend to be defined",
			Subject: &hcl.Range{
				Filename: file,
				Start:    hcl.InitialPos,
				End:      hcl.InitialPos,
			},
		})
		return nil, nil, diags
	}

	ctyTerradude, err := mapToCty(terradude)
	if err != nil {
		return nil, nil, err.(hcl.Diagnostics)
	}

	return configs, ctyTerradude, diags
}
