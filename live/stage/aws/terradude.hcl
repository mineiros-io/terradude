# this file will be created in every aws terradude module (subfolders)
file "terradude-provider-aws.tf" {
  source = "terradude-provider-aws.tf"
}

globals {
  aws_account_id = "0123456789"
  environment    = "stage"
}
