resource "placeos_zone" "p2" {
  name = "Terraform p2"
  parent_id = placeos_zone.terraform_basement.id
  tags = []
}

