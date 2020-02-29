
version = "0.1.0"

terraform {
  module "random_pet" {
    source = "${terradude.base_path}/../modules/random"

    random = [
      { name = "a", data = global.environment },
      { name = "b.b", data = "bb" },
      { name = "c", data = "ccxx" },
      { name = "d", data = "dd" },
    ]
  }

  output "random_pet" {
    value       = module.random_pet
  }
}

globals {
  team = "team superman"
}
