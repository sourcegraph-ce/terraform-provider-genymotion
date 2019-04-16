provider "genymotion" { 
  email       = "myemail@mycompany"
  password    = "mypassword"
}

resource "genymotion_cloud" "Android70" {
  template_uuid = "9f1adea8-e280-460d-9319-580570f61e8c"
  name     = "Android70-ADB-Connected"
  connected_with_adb = true
}

resource "genymotion_cloud" "Android80" {
  template_uuid = "e7a4ecd9-6044-41c7-ace3-ccee5402b590"
  name     = "Android70-ADB-Connected-with-serial-port"
  connected_with_adb = true
  adb_serial_port = "9090"
}

resource "genymotion_cloud" "Android90" {
  template_uuid = "bd402826-4ee6-4598-94df-da4f89021042"
  name     = "Android90-ADB-Not-Connected"
  connected_with_adb = false
}