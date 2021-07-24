terraform {
  required_providers {
    placeos = {
      version = "0.2"
      source  = "hashicorp.com/edu/placeos"
    }
  }
}
resource "placeos_repository" "public_drivers" {
  name = "new public repository"
  folder_name = "terraformfolder2"
  uri = "https://github.com/placeos/drivers"
  repo_type = "driver"
  branch = "master"
}

resource "placeos_driver" "placeos_staff_api" {
  name = "placeos_staff_api"
  file_name = "drivers/place/staff_api.cr"
  description = "Staff api 5"
  role = 1
  module_name = "Staff_API"
  default_uri = "https://nginx"
  repository_id = placeos_repository.public_drivers.id
}

resource "placeos_setting" "place_staff_api_unencrypted" {
  parent_id = placeos_driver.placeos_staff_api.id
  parent_type = "driver"
  encryption_level = 0
  settings_string = "q: 1"
  keys = [
    "q"
  ]
}
