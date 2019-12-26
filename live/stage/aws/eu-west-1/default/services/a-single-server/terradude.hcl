
version = "0.1.0"

terraform {
  module "single-server" {
    source = "${terradude.base_path}/../modules/ec2/modules/single-server/"

    parameter = [""]
    vpc_id = "local.terradude.dependency.vpc.outputs.vpc_id"
    region = global.aws_region
  }
}

globals {
  team = "team superman"
}
