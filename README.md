# terradude - PoC

## Blocks

### `terradude {}`
```hcl
terradude {
  version = "~> 0.1"
}
```
The `terradude` block does not support interpolation of any kind.

#### Attributes
- `version` (required, top-level config) Version constraint for terradude
defining the (minimal) version this repository is compatible with

### `backend "name" {}`
```
backend "s3" {
  bucket         = "terraform-state-files-${global.environment}"
  region         = "eu-west-1"
  dynamodb_table = "terraform-locks"
  key            = "/${global.environment}/${terradude.module_path}/terraform.tfstate"

}
```
The `backend` block supports functions and variable interpolation for
the `global` and `terradude` namespaces.
The `backend` block can be defined only once and should be defined in the top-level
`terradude.hcl` file

### `provider "name" {}`
```
provider "aws" {
  version             = "~> 0.29"
  region              = global.aws_region
  allowed_account_ids = [ global.aws_account_id ]
}
```
The `provider` blocks support functions and variable interpolation for
the `global` and `terradude` namespaces.
The `provider` block can be defined in any level of terradude config.
Multiple definitions of the same provider will not (yet) be merged but cause error.

### `globals {}`
```
globals {
  aws_region      = "eu-west-1"
  aws_account_id  = "0123456789"
  environment     = "production"
  cidr_block      = "10.10.0.0/16"
  subnet_cidrs = [
    cidrsubnet(global.cidr_block, 8, 1),
    cidrsubnet(global.cidr_block, 8, 2),
    cidrsubnet(global.cidr_block, 8, 3)
  ]
}
```
The `globals` block supports functions and variable interpolation of
`global` and `terradude` namespaces.

### `terraform {}`
```
terraform {
  module "ec2" {
    source     = "mineiros-io/ec2-instance/aws"
    version    = "0.1.0"
    subnet_ids = dependency.vpc.private_subnet_ids
  }
}
```
The `terraform` block can only be defined in leaf level `terradude.hcl` files
The `terraform` block allows functions and variable interpolation of
`global`, `dependency` and `terradude` namespaces.

### `dependency "name" {}`
```
dependency "vpc" {
  path = "../vpc"
}
```
The `dependency` block allows variable interpolation of `global`
and `terradude` namespaces.
