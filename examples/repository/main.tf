terraform {
  required_providers {
    placeos = {
      version = "0.2"
      source  = "hashicorp.com/edu/placeos"
    }
  }
}

resource "placeos_repository" "public_drivers" {
  name = "Prueba1"
  folder_name = "terraformfolder2"
  uri = "https://github.com/placeos/drivers"
  repo_type = "driver"
  branch = "master"
}

resource "placeos_driver" "placeos_staff_api" {
  name = "Terraform PlaceStaffApi"
  file_name = "drivers/place/staff_api.cr"
  description = "Staff api 5"
  role = 2
  module_name = "Staff_API"
  default_uri = "https://nginx"
  repository_id = placeos_repository.public_drivers.id
}

resource "placeos_driver" "calendar" {
  name = "Terraform PlaceStaffApi"
  file_name = "drivers/place/calendar.cr"
  description = "Staff api 5"
  role = 2
  module_name = "calendar"
  default_uri = "https://nginx"
  repository_id = placeos_repository.public_drivers.id
}
resource "placeos_driver" "smtp" {
  name = "Terraform PlaceStaffApi"
  file_name = "drivers/place/smtp.cr"
  description = "Staff api 5"
  role = 2
  module_name = "Staff_API"
  default_uri = "https://nginx"
  repository_id = placeos_repository.public_drivers.id
}

resource "placeos_module" "placeos_module_staff_api" {
  custom_name = "TerraformStaffApi"
  uri = "https://nginx"
  driver_id = placeos_driver.placeos_staff_api.id
}

resource "placeos_setting" "driver_setting_staff_api" {
  parent_type = "driver"
  parent_id = placeos_driver.placeos_staff_api.id
  keys = ["name"]
  settings_string = "{\"name\":\"daniel\"}"
  encryption_level = "1"
}


resource "placeos_setting" "module_setting_staff_api" {
  parent_type = "module"
  parent_id = placeos_module.placeos_module_staff_api.id
  keys = ["last_name"]
  settings_string = "{\"last_name\":\"daniel\"}"
  encryption_level = "1"
}

resource "placeos_zone" "terraform_basement" {
  name = "Terraform Basement"
  tags = []
}


resource "placeos_zone" "terraform_p1" {
  name = "Terraform p1"
  parent_id = placeos_zone.terraform_basement.id
  tags = ["p1"]
}