package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

// Create a Genymotion Cloud instance and verify that it gets created with the correct configuration.
func TestAccGenymotionCloudBasicCreate(t *testing.T) {

	var nameBasic = fmt.Sprintf("instance-test-%s", acctest.RandString(10))
	var templateUUIDBasic = "107d757e-463a-4a18-8667-b8dec6e4c87e"
	var checkADB = false

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckGenymotionCloudDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccGenymotionCloudBasic(
					nameBasic,
					templateUUIDBasic,
				),
				Check: resource.ComposeTestCheckFunc(
					testCheckGenymotionCloudInstanceExists("genymotion_cloud.device", nameBasic, templateUUIDBasic, checkADB),
				),
			},
		},
	})
}

// Create a Genymotion Cloud instance with ADB connection and verify that it gets created with the correct configuration.
func TestAccGenymotionCloudBasicWithADBCreate(t *testing.T) {

	var nameBasic = fmt.Sprintf("instance-test-%s", acctest.RandString(10))
	var templateUUIDBasic = "107d757e-463a-4a18-8667-b8dec6e4c87e"
	var checkADB = true

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckGenymotionCloudDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccGenymotionCloudWithADB(
					nameBasic,
					templateUUIDBasic,
				),
				Check: resource.ComposeTestCheckFunc(
					testCheckGenymotionCloudInstanceExists("genymotion_cloud.device", nameBasic, templateUUIDBasic, checkADB),
				),
			},
		},
	})
}

func testCheckGenymotionCloudInstanceExists(resourceName string, name string, templateUUID string, checkADB bool) resource.TestCheckFunc {

	return func(state *terraform.State) error {

		_, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		outputs := state.RootModule().Outputs

		if outputs["template_uuid"].Value != templateUUID {
			return fmt.Errorf(
				`'template_uuid' output is %s; want '%s'`,
				outputs["template_uuid"].Value, templateUUID,
			)
		}

		if outputs["name"].Value != name {
			return fmt.Errorf(
				`'name' output is %s; want '%s'`,
				outputs["name"].Value, name,
			)
		}

		// check instance_uuid is not empty
		if fmt.Sprint(outputs["instance_uuid"].Value) == "" {
			return fmt.Errorf("`instance_uuid` output is empty")
		}

		// check adb serial
		if checkADB {
			if !strings.HasPrefix(fmt.Sprint(outputs["adb_serial"].Value), "localhost:") {
				return fmt.Errorf("`adb_serial` output should start with localhost")
			}
		} else {
			if fmt.Sprint(outputs["adb_serial"].Value) != "0.0.0.0" {
				return fmt.Errorf("`adb_serial` output should be 0.0.0.0")
			}
		}

		return nil
	}
}

func testAccGenymotionCloudBasic(name string, templateUUID string) string {
	return fmt.Sprintf(`
		provider "genymotion" {}

		resource "genymotion_cloud" "device" {
			name		= "%s"
			template_uuid	= "%s"
			connected_with_adb = false
		}
		
		output "template_uuid" {
			value = "${genymotion_cloud.device.template_uuid}"
		}
		output "name" {
			value = "${genymotion_cloud.device.name}"
		}
		output "adb_serial" {
			value = "${genymotion_cloud.device.adb_serial}"
		}
		output "instance_uuid" {
			value = "${genymotion_cloud.device.instance_uuid}"
		}`, name, templateUUID,
	)
}

func testAccGenymotionCloudWithADB(name string, templateUUID string) string {
	return fmt.Sprintf(`
		provider "genymotion" {}

		resource "genymotion_cloud" "device" {
			name		= "%s"
			template_uuid	= "%s"
			connected_with_adb = true
		}
		
		output "template_uuid" {
			value = "${genymotion_cloud.device.template_uuid}"
		}
		output "name" {
			value = "${genymotion_cloud.device.name}"
		}
		output "adb_serial" {
			value = "${genymotion_cloud.device.adb_serial}"
		}
		output "instance_uuid" {
			value = "${genymotion_cloud.device.instance_uuid}"
		}`, name, templateUUID,
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
