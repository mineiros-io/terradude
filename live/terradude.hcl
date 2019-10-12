terradude {
  version = "~> 0.0"
}

backend "s3" {
  bucket          = "terraform-state-${terradude.local.path_elements[0]}"
  key             = "${terradude.local.path}/terraform.tfstate"
  region          = "eu-west-1"
  dynamodb_table  = "terraform-locks"
  encrypt         = true
}

globals {
  team = "teamA"

  aws_region = "eu-west-1"

  aws_account_id_teamA_env   = "12345678901"
  aws_account_id_teamB_prod  = "12345678902"
  aws_account_id_teamB_stage = "12345678903"

  trusted_cidrs_office = ["127.0.0.1/32"]
  trusted_cidrs_vpn_1  = ["127.0.0.2/32"]
  trusted_cidrs_vpn_2  = ["127.0.0.3/32"]
}

define "provider-aws" {
  provider "aws" {
    version             = var.version
    region              = global.aws_region
    allowed_account_ids = var.allowed_account_ids

    assume_role {
      role_arn     = global.assume_role_arn
      session_name = "terradude"
    }
  }
}
