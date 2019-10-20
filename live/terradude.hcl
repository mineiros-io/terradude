terradude {
  version = "~> 0.0"
}

// backend "s3" {
//   bucket          = "terraform-state-${global.environment}"
//   key             = "/${terradude.module_path}/terraform.tfstate"
//   region          = "eu-west-1"
//   dynamodb_table  = "terraform-locks"
//   encrypt         = true
// }

backend "local" {
}

globals {
  team = "teamA"
  path = "abc"
  aws_region = "us-east-1"

  twenty = 5*4

  aws_account_id_teamA_env   = "12345678901"
  aws_account_id_teamB_prod  = "12345678902"
  aws_account_id_teamB_stage = "12345678903"

  trusted_cidrs_office = ["127.0.0.1/32"]
  trusted_cidrs_vpn_1  = ["127.0.0.2/32"]
  trusted_cidrs_vpn_2  = ["127.0.0.3/32"]
}
