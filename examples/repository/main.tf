terraform {
  required_providers {
    placeos = {
      version = "0.2"
      source  = "hashicorp.com/edu/placeos"
    }
  }
}
resource "placeos_repository" "random" {
  name = "new Repository"
  folder_name = "terraformfolder1"
  uri = "https://github.com/placeos/backoffice"
  repo_type = "interface"
}

# Returns all repositories
# output "all_repositories" {
#   value = data.placeos_repositories.all.repositories
# }
