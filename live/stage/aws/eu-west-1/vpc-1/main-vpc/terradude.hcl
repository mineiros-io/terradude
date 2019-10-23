globals {
  cidr = "10.0.0.0/16"
}

terraform {
  module "vpc" {
    source = "terraform-aws-modules/vpc/aws"

    name   = "main-vpc"
    cidr   = global.cidr

    azs = ["eu-west-1a", "eu-west-1b", "eu-west-1c"]
    private_subnets = [
      cidrsubnet(global.cidr, 8, 1),
      cidrsubnet(global.cidr, 8, 2),
      cidrsubnet(global.cidr, 8, 3)
    ]
    public_subnets  = [
      cidrsubnet(global.cidr, 8, 101),
      cidrsubnet(global.cidr, 8, 102),
      cidrsubnet(global.cidr, 8, 103)
    ]

    tags = merge(
      global.default_tags,
      {
        AWSAccount = global.aws_account_id
      }
    )
  }
}
