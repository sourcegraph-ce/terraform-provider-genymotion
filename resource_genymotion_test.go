package main

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

// Create a Genymotion Cloud instance and verify that it gets created with the correct configuration.
func TestAccGenymotionCloudBasicCreate(t *testing.T) {

	var nameBasic = fmt.Sprintf("instance-test-%s", acctest.RandString(10))
	var templateBasic = "Google Nexus 6 - 7.0.0 - API 24 - 1440x2560"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckGenymotionCloudDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccGenymotionCloudBasic(
					nameBasic,
					templateBasic,
				),
				Check: resource.ComposeTestCheckFunc(
					testCheckGenymotionCloudInstanceExists("genymotion_cloud.device", nameBasic, templateBasic),
				),
			},
		},
	})
}

func testCheckGenymotionCloudInstanceExists(resourceName string, name string, template string) resource.TestCheckFunc {

	return func(state *terraform.State) error {

		_, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		outputs := state.RootModule().Outputs

		if outputs["template"].Value != template {
			return fmt.Errorf(
				`'template' output is %s; want '%s'`,
				outputs["template"].Value, template,
			)
		}

		if outputs["name"].Value != name {
			return fmt.Errorf(
				`'name' output is %s; want '%s'`,
				outputs["name"].Value, name,
			)
		}

		// check uuid is not empty
		if fmt.Sprint(outputs["uuid"].Value) == "" {
			return fmt.Errorf("`uuid` output is empty")
		}

		// check adb serial is not empty
		if fmt.Sprint(outputs["adbserial"].Value) == "" {
			return fmt.Errorf("`adbserial` output is empty")
		}

		return nil
	}
}

func testAccGenymotionCloudBasic(name string, template string) string {
	return fmt.Sprintf(`
		provider "genymotion" {}

		resource "genymotion_cloud" "device" {
			name		= "%s"
			template	= "%s"
		}
		
		output "template" {
			value = "${genymotion_cloud.device.template}"
		}
		output "name" {
			value = "${genymotion_cloud.device.name}"
		}
		output "adbserial" {
			value = "${genymotion_cloud.device.adbserial}"
		}
		output "uuid" {
			value = "${genymotion_cloud.device.uuid}"
		}`, name, template,
	)
}

// Check instance specified in the configuration have been destroyed.
func testCheckGenymotionCloudDestroy(state *terraform.State) error {
	for _, res := range state.RootModule().Resources {
		if res.Type != "genymotion_cloud" {
			continue
		}
	}

	return nil
}
