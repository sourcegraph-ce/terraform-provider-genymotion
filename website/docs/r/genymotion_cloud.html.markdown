---
layout: "genymotion"
page_title: "genymotion_cloud"
sidebar_current: "docs-genymotion-resource-genymotion_cloud"
description: |-
  Create a Genymotion Cloud Android virtual device
---

# Genymotion Cloud instance

 genymotion_cloud resource provides a [Genymotion Cloud Android virtual device](https://cloud.geny.io?&utm_source=web-referral&utm_medium=docs&utm_campaign=terraform&utm_content=signup).

## Example Usage

```hcl
# Create an Genymotion Cloud SaaS Android instance
resource "genymotion_cloud" "Android90" {
  recipe_uuid = "143eb44a-1d3a-4f27-bcac-3c40124e2836"
  name     = "Android90"
  adbconnect = true
  adb_serial_port = "9090"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the Android instance.
* `recipe_uuid` - (Required) Recipe UUID is the identifier used when starting an instance. It can be retrieved using `gmsaas recipes list`
* `adbconnect` - (Optional) If is true, it will connect the instance to ADB. Defaults to "true".
* `adb_serial_port` - (Optional) If the --adb_serial_port <PORT> option is set, the instance will be connected to ADB on localhost:<PORT>.
