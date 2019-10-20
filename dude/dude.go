package dude

import (
	"os"
	"io/ioutil"
	"path/filepath"
	"github.com/zclconf/go-cty/cty"
	"github.com/mineiros-io/terradude/config"
	"github.com/rs/zerolog/log"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func RunFmt(file string) error {
	log.Info().Msgf("> start processing %v", filepath.Dir(file))

	hclconfigs, terradude, diags := config.LoadConfigs(file)
	if diags.HasErrors() {
		log.Fatal().Msg(diags.Error())
	}

	globals, diags := config.DecodeGlobalCty(hclconfigs)
	if diags.HasErrors() {
		log.Fatal().Msg(diags.Error())
	}

	ctx := &hcl.EvalContext{}
	ctx.Variables = map[string]cty.Value{}
	ctx.Variables["global"] = *globals
	ctx.Variables["terradude"] = *terradude

	backend, diags := config.DecodeBackendBlock(hclconfigs, ctx)
	if diags.HasErrors() {
		log.Fatal().Msg(diags.Error())
	}

	terraform, diags := config.DecodeTerraformBlock(hclconfigs, ctx)
	if diags.HasErrors() {
		log.Fatal().Msg(diags.Error())
	}

	providers, diags := config.DecodeProviderBlocks(hclconfigs, ctx)
	if diags.HasErrors() {
		log.Fatal().Msg(diags.Error())
	}

	f := hclwrite.NewEmptyFile()

	log.Info().Msgf("+ appending backend config for %s", filepath.Dir(file))
	f.Body().AppendBlock(backend)
	log.Info().Msgf("+ appending provider config for %s", filepath.Dir(file))
	for _, provider := range providers {
		f.Body().AppendNewline()
		f.Body().AppendBlock(provider)
	}
	log.Info().Msgf("+ appending terraform config for %s", filepath.Dir(file))
	f.Body().AppendNewline()
	f.Body().AppendBlock(terraform)
	log.Info().Msgf("= rendered config for %s", filepath.Dir(file))

	config := terradude.AsValueMap()
	err := os.MkdirAll(config["terraform_path"].AsString(), 0755)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	terradudeTF := config["terraform_path"].AsString() + "/terradude.tf"
	err = ioutil.WriteFile(terradudeTF, f.Bytes(), 0644)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	log.Info().Msgf("= saved config in %s", filepath.Dir(terradudeTF))
	log.Info().Msgf("< finished processing %v", filepath.Dir(file))
	return nil
}
