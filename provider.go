package main

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/hashicorp/go-multierror"
    "log"
    "fmt"
    "os/exec"
)

func Provider() *schema.Provider {
    return &schema.Provider{
        Schema: map[string]*schema.Schema{
            "email": &schema.Schema{
    	        Type:     schema.TypeString,
                Required: true,
                DefaultFunc: schema.EnvDefaultFunc("GENYMOTION_EMAIL", nil),
                Description: "Email for the Genymotion Cloud account",
			},
			
			"password": &schema.Schema{
	            Type:     schema.TypeString,
                Required: true,
                DefaultFunc: schema.EnvDefaultFunc("GENYMOTION_PASSWORD", nil),
				Description: "Password for the Genymotion Cloud account",
			},
			"license_key": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
                DefaultFunc: schema.EnvDefaultFunc("GENYMOTION_LICENSE_KEY", nil),
				Description: "License key for the Genymotion Cloud account",
			},

        },

        ResourcesMap: map[string]*schema.Resource{
            "genymotion_cloud": resourceGenymotion(),
        },

        ConfigureFunc: providerConfigure,
    }
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

    config := GenymotionConfig{
        Email      : d.Get("email").(string),
        Password   : d.Get("password").(string),
        LicenseKey : d.Get("license_key").(string),
    }

    if err := config.validate(); err != nil {
        return nil, err
    }

    // Register Genymotion Account
    log.Println("[INFO] Register Genymotion Account")
	cmd_login := exec.Command(
        "gmtool", "config", "username", config.Email, "password", config.Password)
    output, err_login := cmd_login.CombinedOutput()
    if err_login != nil {
        return fmt.Errorf("Error: %s", output), nil
    } else {
        fmt.Println(string(output))
    }

    // Register Genymotion License key
    log.Println("[INFO] Register Genymotion License key")
	cmd_register := exec.Command(
        "gmtool", "license", "register", config.LicenseKey)
    output, err_register := cmd_register.CombinedOutput()
    if err_register != nil {
        return fmt.Errorf("Error: %s", output), nil
    } else {
        fmt.Println(string(output))
    }

    return nil, err_login
}


func (c GenymotionConfig) validate() error {
    var err *multierror.Error

    if c.Email == "" {
        err = multierror.Append(err, fmt.Errorf("Email must be configured for the Genymotion Cloud Provider"))
    }
    if c.Password == "" {
        err = multierror.Append(err, fmt.Errorf("Password must be configured for the Genymotion Cloud Provider"))
    }
    if c.LicenseKey == "" {
        err = multierror.Append(err, fmt.Errorf("License key must be configured for the Genymotion Cloud Provider"))
    }

    return err.ErrorOrNil()
}

type GenymotionConfig  struct {
    Email       string
    Password    string
    LicenseKey string
}
