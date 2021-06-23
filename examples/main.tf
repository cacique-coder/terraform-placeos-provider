terraform {
  required_providers {
    placeos = {
      version = "0.2"
      source  = "hashicorp.com/edu/placeos"
    }
  }
}

provider "placeos" {}

module "psl" {
  source = "./repository"

  repository_name = "wcpdrivers"
}

output "psl" {
  value = module.psl.repository
}
output "repositories" {
  value = module.psl.all_repositories
}
