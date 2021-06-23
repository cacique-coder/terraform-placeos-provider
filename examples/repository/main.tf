terraform {
  required_providers {
    placeos = {
      version = "0.2"
      source  = "hashicorp.com/edu/placeos"
    }
  }
}

variable "repository_name" {
  type    = string
  default = "Vagrante espresso"
}

data "placeos_repository" "all" {

}

# Returns all repositories
output "all_repositories" {
  value = data.placeos_repository.all.repositories
}

# Only returns packer spiced latte
output "repository" {
  value = {
    for repository in data.placeos_repository.all.repositories :
    repository.id => repository
    if repository.name == var.repository_name
  }
}
