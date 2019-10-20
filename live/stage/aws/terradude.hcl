provider "aws" {
  version             = "~> 2.7"
  allowed_account_ids = [ global.aws_account_id ]
  region              = global.aws_region
}
