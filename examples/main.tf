terraform {
  required_providers {
    placeos = {
      version = "0.2"
    source  = "hashicorp.com/edu/placeos"
    }
  }
}

provider "placeos" {
  username = var.username
	password =  var.password
	host =  var.host
  client_id = var.client_id
  client_secret = var.client_secret
  insecure_ssl = var.insecure_ssl
}

module "psl" {
  source = "./repository"

}
output "repositories" {
  value = module.psl.all_repositories
}
