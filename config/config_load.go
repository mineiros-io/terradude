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

	configFileName, _ = filepath.Abs(configFileName)

	files     := util.SearchUp(configFileName, DefaultConfigFileBaseName)
	backend   := false
	terradude := map[string]cty.Value{}

	for _, file = range files {
		var config Config
		abs, err := filepath.Abs(".")
		rel, err := filepath.Rel(abs, file)
		if err != nil {
			panic(err)
		}
		log.Debug().Msgf("  including config %s", rel)

		err = hclsimple.DecodeFile(file, nil, &config)
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
			log.Debug().Msgf("    found terradude.version in %s - stop including config files", file)

			dir    := filepath.Dir(file)
			rel, _ := filepath.Rel(dir, configFileName)
			abs, _ := filepath.Abs(dir)
			mod    := filepath.Dir(rel)

			terradude["module_path"]        = cty.StringVal(mod)
			terradude["base_path"]          = cty.StringVal(dir)
			terradude["absolute_base_path"] = cty.StringVal(abs)
			terradude["terraform_path"]     = cty.StringVal(abs + "/.terradude/" + mod)
			log.Debug().Msgf("    setting terradude.module_path        = %#v", terradude["module_path"].AsString())
			log.Debug().Msgf("    setting terradude.base_path          = %#v", terradude["base_path"].AsString())
			log.Debug().Msgf("    setting terradude.absolute_base_path = %#v", terradude["absolute_base_path"].AsString())
			log.Debug().Msgf("    setting terradude.terraform_path     = %#v", terradude["terraform_path"].AsString())
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
