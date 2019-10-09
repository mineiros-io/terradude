provider "aws" {
  version             = "~> 0.29"
  region              = global.aws_region
  allowed_account_ids = [ global.aws_account_id ]

  assume_role {
    role_arn     = global.assume_role_arn
    session_name = "terradude"
  }
}
