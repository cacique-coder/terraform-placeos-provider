terraform {
  required_providers {
    placeos = {
      version = "0.2"
    source  = "hashicorp.com/edu/placeos"
    }
  }
}

provider "placeos" {
  username = "support@place.tech"
	password =  "development"
	host =  "https://localhost:8443"
  client_id = "b52e653071c45353dbff4e8f47d51cdf"
  client_secret = "288ix5hR8y_lyoNeD0ujd9w35FKjr1FAQySY2CbQ-PDB3tRY1ECZ3w"

  # insecureSsl = false
}

module "psl" {
  source = "./repository"

}
output "repositories" {
  value = module.psl.all_repositories
}
