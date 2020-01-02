variable "random" {
  type = list(any)
}

locals {
  random = { for r in var.random : lookup(r, "id", lower(r.name)) => r }
}

resource "random_pet" "count" {
  count  = length(var.random)
  prefix = var.random[count.index].name
}

resource "random_pet" "for_each" {
  for_each = local.random
  prefix   = each.value.name
}

output "for_each" {
  value = random_pet.for_each
}

output "count" {
  value = random_pet.count
}
