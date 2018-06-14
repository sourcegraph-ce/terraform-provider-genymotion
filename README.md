# terraform-provider-genymotion

[Terraform](https://www.terraform.io) Custom Provider for [Genymotion Cloud SAAS](https://www.genymotion.com/cloud/)

## Description

This is a custom terraform provider for managing Android devices within the Genymotion Cloud SAAS platform.

## Requirements

* [hashicorp/terraform](https://github.com/hashicorp/terraform)
* [Go](https://golang.org/doc/install) 1.9 (to build the provider plugin)
* [Genymotion Cloud Account](https://www.genymotion.com/acount/create/)
* [Genymotion Desktop](https://www.genymotion.com/download/) (**gmtool** binary must be set in the PATH)
* [Genymotion Cloud License](https://www.genymotion.com/pricing-and-licensing/)

## Installation

1. Download the appropriate release for your system: https://github.com/genymobile/terraform-provider-genymotion/releases

1. Unzip/untar the archive.

1. Move it into `$HOME/.terraform.d/plugins`:

    ```sh
    $ mkdir -p $HOME/.terraform.d/plugins
    $ mv terraform-provider-genymotion $HOME/.terraform.d/plugins/terraform-provider-genymotion
    ```

1. After placing it into your plugins directory, run `terraform init` to initialize it.

  This will find the plugin locally.


## Using the provider

### Setup ###

The provider takes configuration arguments for setting up your Genymotion account within Terraform. The following example shows you how to explicitly configure the provider using your account information.

```hcl
provider "genymotion" {
  email = ""
  password  = ""
  license_key = ""
}
```

The following arguments are supported.

- `email` - (Required) This is the email of the Genymotion account. It can also be provided via the `GENYMOTION_EMAIL` environment variable.
- `password` - (Required) This is the password of the Genymotion account. It can also be provided via the `GENYMOTION_PASSWORD` environment variable.
- `license_key` - (Required )This is the license key of the Genymotion account. It can also be provided via the `GENYMOTION_LICENSE_KEY` environment variable.

### Example ###

The following example shows you how to configure a LX branded zone running Ubuntu.

```hcl
provider "genymotion" {
    email = "name@company.com"
    password = "its@wEsOme"
    license_key = "1234-1234-5678-5678"
}  
```

If you are using environnement variables, just add the provider : 
```hcl
provider "genymotion' {}
```


## Resources Providers ##

### Example - One device ###

The following example shows you how to start one device on Genymotion Cloud SAAS platform.

```hcl
# use env vars to configure the provider
provider "genymotion" {}


resource "genymotion_cloud" "myAndroid70" {
    template = "Google Nexus 6 - 7.0.0 - API 24 - 1440x2560"
    name     = "myAndroidDevice70"
}
```

### Example - Several devices ###

The following example shows you how to start several devices on Genymotion Cloud SAAS platform.

```hcl
# use env vars to configure the provider
provider "genymotion" {}

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
```

### Example - Start several devices based on the same template

The following example shows you how to start 3 devices based on the same template.

```hcl
# use env vars to configure the provider
provider "genymotion" {}

resource "genymotion_cloud" "Android70" {
  template = "Google Nexus 6 - 7.0.0 - API 24 - 1440x2560"
  name     = "MyAndroid70"

  count = 3
}
```

## Building The Provider

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-genymotion`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:Genymobile/terraform-provider-genymotion
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-genymotion
$ make build
```

Initialize your Terraform project by passing in the directory that contains your custom built provider binary, `terraform-provider-genymotion`. This is typically `$GOPATH/bin`.

```sh
$ terraform version
Terraform v0.17.0

$ terraform init --plugin-dir=$GOPATH/bin
```