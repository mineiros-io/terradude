
version = "0.1.0"

terraform {
  module "single-server" {
    source = "../../../../../../../modules/ec2/modules/single-server/"

    parameter = [""]
    vpc_id = local.terradude.dependency.vpc.outputs.vpc_id
    region = local.terradude.global.aws_region
  }
}

globals {
  team = "team superman"
}
