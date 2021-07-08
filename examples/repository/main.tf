terraform {
  required_providers {
    placeos = {
      version = "0.2"
      source  = "hashicorp.com/edu/placeos"
    }
  }
}
data "placeos_repository" "all" {
}

# Returns all repositories
output "all_repositories" {
  value = data.placeos_repository.all.repositories
}
