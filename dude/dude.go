package dude

import (
	"github.com/mineiros-io/terradude/config"
	"github.com/rs/zerolog/log"
)

func RunFmt(file string) error {

	hclconfigs, diags := config.LoadConfigs(file)
	if diags.HasErrors() {
		log.Fatal().Msg(diags.Error())
	}

	terraform, diags := config.DecodeTerraformBlock(hclconfigs)
	if diags.HasErrors() {
		log.Fatal().Msg(diags.Error())
	}

	globals, diags := config.DecodeGlobalValues(hclconfigs)
	if diags.HasErrors() {
		log.Fatal().Msg(diags.Error())
	}

	backend, diags := config.DecodeBackendBlock(hclconfigs, globals)
	if diags.HasErrors() {
		log.Fatal().Msg(diags.Error())
	}

	log.Debug().Msgf("%#v", backend)
	log.Debug().Msgf("%#v", terraform)

	log.Printf("### Final globals")
  for _, g := range globals {
		log.Printf("  global.%s = %#v", g.Name, g.Value)
	}

	return nil
}
