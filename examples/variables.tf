variable "username" {
  type = string
  default = "support@place.tech"
}

variable "password" {
  type = string
}

variable "host" {
  type = string
  default = "https://localhost:8443"
}

variable "client_id" {
  type = string
}

variable "client_secret" {
  type = string
}

variable "insecure_ssl" {
  type = bool
  default = true
}