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


resource "placeos_module" "placeos_module_staff_api" {
  name = "placeos_staff_apiX"
  custom_name = "Staff_API"
  uri = "https://nginx"
  driver_id = placeos_driver.placeos_staff_api.id
}
