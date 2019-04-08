package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"email": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GENYMOTION_EMAIL", nil),
				Description: "Email for the Genymotion Cloud account",
			},

			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GENYMOTION_PASSWORD", nil),
				Description: "Password for the Genymotion Cloud account",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"genymotion_cloud": resourceGenymotion(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	var err error

	config := GenymotionConfig{
		Email:      d.Get("email").(string),
		Password:   d.Get("password").(string)
	}

	// Check mandatory fields
	if err := config.validate(); err != nil {
		return nil, err
	}

	// Connect to Genymotion account
	if err := config.connect(); err != nil {
		return nil, err
	}

	return nil, err
}

func (c GenymotionConfig) validate() error {
	var err *multierror.Error

	if c.Email == "" {
		err = multierror.Append(err, fmt.Errorf("Email must be configured for the Genymotion Cloud Provider"))
	}
	if c.Password == "" {
		err = multierror.Append(err, fmt.Errorf("Password must be configured for the Genymotion Cloud Provider"))
	}

	return err.ErrorOrNil()
}

func (c GenymotionConfig) connect() error {
	// Register Genymotion Account
	log.Println("[INFO] Register Genymotion Account")
	cmd := exec.Command(
		"gmtool", "config", "username", c.Email, "password", c.Password)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Error: %s", output)
	} else {
		fmt.Println(string(output))
	}

	return nil
}

type GenymotionConfig struct {
	Email      string
	Password   string
}
