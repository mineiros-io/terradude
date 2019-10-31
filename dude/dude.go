package dude

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	tflang "github.com/hashicorp/terraform/lang"
	"github.com/mineiros-io/terradude/config"
	"github.com/rs/zerolog/log"
	"github.com/zclconf/go-cty/cty"
	"io/ioutil"
	"os"
	"path/filepath"
)

func RunFmt(file string) error {
	log.Info().
		Str("module", filepath.Dir(file)).
		Msg("processing module")

	hclconfigs, terradude, diags := config.LoadConfigs(file)
	if diags.HasErrors() {
		log.Fatal().
			Str("file", file).
			Str("module", filepath.Dir(file)).
			Err(diags).
			Msg("could not load config file")
	}

	tfscope := tflang.Scope{
		BaseDir: filepath.Dir(file),
	}

	ctx := &hcl.EvalContext{}
	ctx.Variables = map[string]cty.Value{}
	ctx.Functions = tfscope.Functions()
	ctx.Variables["terradude"] = *terradude

	globals, diags := config.DecodeGlobalCty(hclconfigs, ctx)
	if diags.HasErrors() {
		log.Fatal().
			Str("file", file).
			Str("module", filepath.Dir(file)).
			Err(diags).
			Msg("could not decode globals block")
	}

	ctx.Variables["global"] = *globals

	backend, diags := config.DecodeBackendBlock(hclconfigs, ctx)
	if diags.HasErrors() {
		log.Fatal().
			Str("file", file).
			Str("module", filepath.Dir(file)).
			Err(diags).
			Msg("could not decode backend block")
	}

	terraform, diags := config.DecodeTerraformBlock(hclconfigs, ctx)
	if diags.HasErrors() {
		log.Fatal().
			Str("file", file).
			Str("module", filepath.Dir(file)).
			Err(diags).
			Msg("could not decode terraform block")
	}

	providers, diags := config.DecodeProviderBlocks(hclconfigs, ctx)
	if diags.HasErrors() {
		log.Fatal().
			Str("file", file).
			Str("module", filepath.Dir(file)).
			Err(diags).
			Msg("could not provider blocks")
	}

	f := hclwrite.NewEmptyFile()

	log.Info().
		Str("module", filepath.Dir(file)).
		Msg("appending backend config")
	f.Body().AppendBlock(backend)
	log.Info().
		Str("module", filepath.Dir(file)).
		Msg("appending provider config")
	for _, provider := range providers {
		f.Body().AppendNewline()
		f.Body().AppendBlock(provider)
	}
	log.Info().
		Str("module", filepath.Dir(file)).
		Msg("appending terraform module config")
	f.Body().AppendNewline()
	f.Body().AppendBlock(terraform)
	log.Info().
		Str("module", filepath.Dir(file)).
		Msg("completed rendering config")

	config := terradude.AsValueMap()
	err := os.MkdirAll(config["terraform_path"].AsString(), 0755)
	if err != nil {
		log.Fatal().
			Str("module", filepath.Dir(file)).
			Str("path", config["terraform_path"].AsString()).
			Err(err).
			Msg("could not create directory structure")
	}

	terradudeTF := config["terraform_path"].AsString() + "/terradude.tf"
	err = ioutil.WriteFile(terradudeTF, f.Bytes(), 0644)
	if err != nil {
		log.Fatal().
			Str("module", filepath.Dir(file)).
			Str("path", config["terraform_path"].AsString()+"/terradude.tf").
			Err(err).
			Msg("could not create terraform file")
	}
	log.Info().
		Str("module", filepath.Dir(file)).
		Str("path", config["terraform_path"].AsString()+"/terradude.tf").
		Msg("created terraform file")
	return nil
}
