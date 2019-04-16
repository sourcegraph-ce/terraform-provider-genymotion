# terraform-provider-genymotion

[Terraform](https://www.terraform.io) Custom Provider for [Genymotion Cloud SAAS](https://www.genymotion.com/cloud/)

## Description

This is a custom terraform provider for managing Android devices within the Genymotion Cloud SAAS platform.

## Requirements

* [hashicorp/terraform](https://github.com/hashicorp/terraform)
* [Go](https://golang.org/doc/install) 1.9 (to build the provider plugin)
* [Genymotion Cloud Account](https://cloud.geny.io/)
* [gmsaas] (**pip install gmsaas**)
* [Android SDK](https://developer.android.com/studio/index.html#downloads)


## gmsaas

You must install gmsaas command line tool using `pip install gmsaas`

You also must to configure android sdk path using `gmsaas config set android-sdk-path <PATH_TO_ANDROID_SDK>`

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

The provider takes configuration arguments for setting up your Genymotion Cloud account within Terraform. The following example shows you how to explicitly configure the provider using your account information.

```hcl
provider "genymotion" {
  email = ""
  password  = ""

}
```

The following arguments are supported.

- `email` - (Required) This is the email of the Genymotion Cloud account. It can also be provided via the `GENYMOTION_EMAIL` environment variable.
- `password` - (Required) This is the password of the Genymotion Cloud account. It can also be provided via the `GENYMOTION_PASSWORD` environment variable.

### Example ###

The following example shows you how to configure Genymotion Cloud provider.

```hcl
provider "genymotion" {
    email = "name@company.com"
    password = "its@wEsOme"
}  
```

If you use environnement variables, just add the provider : 
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
    template_uuid = "a0a9c90a-b391-42f4-b77b-ae0561d74bbe"
    name     = "myAndroidDevice70"
}
```

### Example - Several devices ###

The following example shows you how to start several devices on Genymotion Cloud SAAS platform.

```hcl
# use env vars to configure the provider
provider "genymotion" {}

resource "genymotion_cloud" "Android70" {
  template_uuid = "a0a9c90a-b391-42f4-b77b-ae0561d74bbe"
  name     = "MyAndroid70"
}

resource "genymotion_cloud" "Android71" {
  template_uuid = "80a67ae9-430c-4824-a386-befbb19518b9"
  name     = "MyAndroid71"
}

resource "genymotion_cloud" "Android80" {
  template_uuid = "a59951f2-ed13-40f9-80b9-3ddceb3c89f5"
  name     = "MyAndroid80"
}
```

### Example - Start several devices based on the same template

The following example shows you how to start 3 devices based on the same template.

```hcl
# use env vars to configure the provider
provider "genymotion" {}

resource "genymotion_cloud" "device" {
  count = 3

  template_uuid = "a0a9c90a-b391-42f4-b77b-ae0561d74bbe"
  name     = "MyAndroid70-${count.index}"  
}
```

### Example - Start one device with adb connection

The following example shows you how to start 1 device and connect it with adb.
* By default connected_with_adb is equal to true, so the paramater is optionnal
* if you don"t want to connect the device with adb, set `connected_with_adb = false`

```hcl
# use env vars to configure the provider
provider "genymotion" {}

resource "genymotion_cloud" "device" {

  template = "a0a9c90a-b391-42f4-b77b-ae0561d74bbe"
  name     = "MyAndroid70"
  connected_with_adb = true
  
}
```

### Example - Start one device with adb connection and specify an adb serial port

The following example shows you how to start 1 device and connect it with adb specify an adb serial port ( e.g localhost:7090)

```hcl
# use env vars to configure the provider
provider "genymotion" {}

resource "genymotion_cloud" "device" {

  template = "a0a9c90a-b391-42f4-b77b-ae0561d74bbe"
  name     = "MyAndroid70"
  adb_serial_port = "7090"
  
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
Terraform v0.11.13

$ terraform init --plugin-dir=$GOPATH/bin
```

## Test the provider

Run acceptance tests
```sh
$ export GENYMOTION_EMAIL={GENYMOTION_ACCOUNT}
$ export GENYMOTION_PASSWORD={GENYMOTION_PASSWORD}
$ make testacc
```

or 

Build and test 
```sh
$ export GENYMOTION_EMAIL={GENYMOTION_ACCOUNT}
$ export GENYMOTION_PASSWORD={GENYMOTION_PASSWORD}
$ make
```

## Release the provider

```sh
$ export GENYMOTION_EMAIL={GENYMOTION_ACCOUNT}
$ export GENYMOTION_PASSWORD={GENYMOTION_PASSWORD}
$ make all
```

Copy the generated binaries to github releases.