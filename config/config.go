package config

import (
	"github.com/hashicorp/hcl/v2"
)

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
  Remain hcl.Body `hcl:",remain"`
}

type Terraform struct {
  Remain hcl.Body `hcl:",remain"`
}

type Dependency struct {
	Name string `hcl:"name,label"`
  Remain hcl.Body `hcl:",remain"`
}

type Backend struct {
	Name string `hcl:"name,label"`
  Remain hcl.Body `hcl:",remain"`
}

type Globals struct {
  Remain hcl.Body `hcl:",remain"`
}

type Define struct {
	Name string `hcl:"name,label"`
  Remain hcl.Body `hcl:",remain"`
}

type Template struct {
	Name string `hcl:"name,label"`
  Body hcl.Body `hcl:",remain"`
}
