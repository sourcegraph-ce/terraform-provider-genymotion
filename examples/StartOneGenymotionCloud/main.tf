provider "genymotion" {
}

variable "template_uuid" {
  type    = "string"
  default = "107d757e-463a-4a18-8667-b8dec6e4c87e"
}

resource "genymotion_cloud" "Android70" {
  template_uuid = "${var.template_uuid}"
  name     = "DeviceConnectedWithAdb"
  connected_with_adb = true
  adb_serial_port = "7090"
}
