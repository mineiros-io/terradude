package config

import (
	"github.com/hashicorp/hcl/v2"
)

const DefaultConfigFileBaseName = "terradude.hcl"

type Config struct {
	Terradude  *Terradude    `hcl:"terradude,block"`
	Terraform  *Terraform    `hcl:"terraform,block"`
	Dependency []*Dependency `hcl:"dependency,block"`
	Backend    *Backend      `hcl:"backend,block"`
	Provider   []*Provider   `hcl:"provider,block"`
	Globals    *Globals      `hcl:"globals,block"`
	Remain     hcl.Body      `hcl:",remain"`
}

type Terradude struct {
	Version string   `hcl:"version"`
	Remain  hcl.Body `hcl:",remain"`
}

type Terraform struct {
	Module Module `hcl:"module,block"`
}

type Module struct {
	Name   string   `hcl:"name,label"`
	Source string   `hcl:"source,attr"`
	Body   hcl.Body `hcl:",remain"`
}

type Dependency struct {
	Name string   `hcl:"name,label"`
	Path string   `hcl:"path,attr"`
	Body hcl.Body `hcl:",remain"`
}

type Backend struct {
	Name string   `hcl:"name,label"`
	Body hcl.Body `hcl:",remain"`
}

type Provider struct {
	Name    string   `hcl:"name,label"`
	Version *string  `hcl:"version,attr"`
	Alias   *string  `hcl:"alias,attr"`
	Body    hcl.Body `hcl:",remain"`
}

type Globals struct {
	Body hcl.Body `hcl:",remain"`
}
