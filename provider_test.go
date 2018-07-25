package main

import (
	"testing"
	"os"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider


func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]terraform.ResourceProvider{
		"genymotion": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("GENYMOTION_EMAIL"); v == "" {
		t.Fatal("GENYMOTION_EMAIL must be set for acceptance tests")
	}
	if v := os.Getenv("GENYMOTION_PASSWORD"); v == "" {
		t.Fatal("GENYMOTION_PASSWORD must be set for acceptance tests")
	}
	if v := os.Getenv("GENYMOTION_LICENSE_KEY"); v == "" {
		t.Fatal("GENYMOTION_LICENSE_KEY must be set for acceptance tests")
	}
}