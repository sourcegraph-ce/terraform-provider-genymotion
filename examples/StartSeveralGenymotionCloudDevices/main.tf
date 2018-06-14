provider "genymotion" {
  email       = "myemail@mycompany"
  password    = "mypassword"
  license_key = "mylicense_key"
}

resource "genymotion_cloud" "Android70" {
  template = "Google Nexus 6 - 7.0.0 - API 24 - 1440x2560"
  name     = "MyAndroid70"
}

resource "genymotion_cloud" "Android71" {
  template = "Google Nexus 6 - 7.1.0 - API 25 - 1440x2560"
  name     = "MyAndroid71"
}

resource "genymotion_cloud" "Android80" {
  template = "Google Nexus 6 - 8.0 - API 26 - 1440x2560"
  name     = "MyAndroid80"
}
