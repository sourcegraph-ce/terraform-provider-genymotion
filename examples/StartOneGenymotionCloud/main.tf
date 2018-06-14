provider "genymotion" {}

variable "template" {
  type    = "string"
  default = "Google Nexus 6 - 7.0.0 - API 24 - 1440x2560"
}

variable "name" {
  type    = "string"
  default = "MyTestDevice"
}

resource "genymotion_cloud" "Android70" {
  template = "${var.template}"
  name     = "${var.name}"
}
