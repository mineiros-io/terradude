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
  default_tags = {
    Terradude       = "true"
    TerradudeModule = terradude.module_path
    Terraform       = "true"
    Environment     = split("/", terradude.module_path)[0]
  }

  aws_region = "us-east-1"

  aws_account_id = lookup({
      stage      = "12345678901"
      production = "12345678902"
    },
    split("/", terradude.module_path)[0],
    "none"
  )
}
