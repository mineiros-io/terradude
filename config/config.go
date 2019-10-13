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
  Remain 		 hcl.Body      `hcl:",remain"`
}

type Terradude struct {
	Version string   `hcl:"version"`
  Remain  hcl.Body `hcl:",remain"`
}

type Terraform struct {
  Remain hcl.Body `hcl:",remain"`
}

type Dependency struct {
	Name   string   `hcl:"name,label"`
	Path   string   `hcl:"path,attr"`
  Remain hcl.Body `hcl:",remain"`
}

type Backend struct {
	Name   string   `hcl:"name,label"`
  Body   hcl.Body `hcl:",remain"`
}

type BackendS3 struct {
	Bucket             string   `hcl:"bucket"`
	Key                string   `hcl:"key"`
	Region             string   `hcl:"region"`
	dynamodb_endpoint  *string  `hcl:"dynamodb_endpoint"`
	// endpoint
	// iam_endpoint
	// sts_endpoint
	// encrypt
	// acl
	// access_key
	// secret_key
	// kms_key_id
	// lock_table
	// dynamodb_table
	// profile
	// shared_credentials_file
	// token
	// skip_credentials_validation
	// skip_get_ec2_platforms
	// skip_region_validation
	// skip_requesting_account_id
	// skip_metadata_api_check
	// role_arnrole_arn
	// session_name
	// external_id
	// assume_role_policy
	// workspace_key_prefix
	// force_path_style
	// max_retries
	//
	//
	//
}


type Provider struct {
	Name string   `hcl:"name,label"`
  Body hcl.Body `hcl:",remain"`
}

type Globals struct {
  Body hcl.Body `hcl:",remain"`
}
