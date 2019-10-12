package config

const DefaultConfigPath = "terradude.hcl"

// represents a parsed config
type Config struct {
	Terradude  *Terradude  `hcl:"terradude,block"`
	Terraform  *Terraform  `hcl:"terraform,block"`
	Dependency *Dependency `hcl:"dependency,block"`
	Backend    *Backend    `hcl:"backend,block"`
	Globals    *Globals    `hcl:"globals,block"`
	Define     *Define     `hcl:"define,block"`
	Template   *Template   `hcl:"template,block"`
}

type Terradude struct {
	Version string `hcl:"version"`
}

type Terraform struct {
}

type Dependency struct {
}

type Backend struct {
}

type Globals struct {
}

type Define struct {
}

type Template struct {
}
