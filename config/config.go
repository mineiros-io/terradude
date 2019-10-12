package config

const DefaultConfigPath = "terradude.hcl"

// represents a parsed config
type Config struct {
	Terradude *Terradude `hcl:"terradude,block"`
	Backend   *Backend   `hcl:"backend,block"`
	Globals   *Globals   `hcl:"globals,block"`
}

type Terradude struct {
	Version string `hcl:"version"`
}

type Backend struct {
}

type Globals struct {
}
