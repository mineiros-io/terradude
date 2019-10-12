template "provider-aws" {
  version = "~> 0.29"

  allowed_account_ids = [ "0123456789" ]
}

globals {
  environment = "stage"
}
