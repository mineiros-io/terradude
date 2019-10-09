package config

const DefaultConfigPath = "terradude.hcl"

// represents a parsed config
type Config struct {
	Terradude *Terradude `hcl:"terradude,block"`
	Backend   *Backend   `hcl:"backend"`
	Globals   *Globals   `hcl:"globals"`
}

type Terradude struct {
	Version string `hcl:"version"`
}

type Backend struct {
}

type Globals struct {
}
