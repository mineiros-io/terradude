dependency "vpc" {
  path    = "../main-vpc"
}

dependency "vpc-peering" {
  path    = "../vpc-peering"
}

terraform {
  module "single-server" {
    source  = "terraform-aws-modules/ec2-instance/aws"
    version = "~> 2.0"

    name = "main-ec2"

    root_block_device = [
      {
        volume_type = "gp2"
        volume_size = 10
      }
    ]

    instance_type = "t2.small"
    ami           = "ami-ebd02392"

    subnet_idsubnet_ids = dependency.vpc.outputs.public_subnets

    tags = {
      Terradude = "true"
      Terraform = "true"
      Environment = global.environment
    }
  }
}
