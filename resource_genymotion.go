package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"bufio"
)

func resourceGenymotion() *schema.Resource {
    return &schema.Resource{
		Create: resourceGenymotionCreate,
		Read: resourceGenymotionRead,
        Delete: resourceGenymotionDelete,

    	Schema: map[string]*schema.Schema{
	        "template": &schema.Schema{
    	        Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			
			"name": &schema.Schema{
	            Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"uuid": &schema.Schema{
	            Type:     schema.TypeString,
				Computed: true,
			},
			"adbserial": &schema.Schema{
	            Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGenymotionCreate(d *schema.ResourceData, m interface{}) error {
	template := d.Get("template").(string)
	name := d.Get("name").(string)

	 // Start Genymotion Cloud Device
	cmd := exec.Command("gmtool", "admin", "--cloud", "startdisposable", template, name )
	output, err := cmd.CombinedOutput()
	if err != nil {
    	return fmt.Errorf("Error: %s", output)
	} else {
    	fmt.Println(string(output))
	}

    id := d.Get("name").(string)
	d.SetId(id)
	
    return resourceGenymotionRead(d, m)
}

func resourceGenymotionRead(d *schema.ResourceData, m interface{}) error {

	name := d.Get("name").(string)
 
	// Retrieve genymotion Cloud device informations
	admin_list := exec.Command("gmtool", "--cloud", "admin", "list")
    stdout, err := admin_list.StdoutPipe()
    if err != nil { return nil}
    admin_list.Start()
    scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
        line := scanner.Text()
        s := strings.Split(line, "|")
        if len(s) >= 4 {
            actual_name := strings.Trim(s[3], " ")
            if strings.EqualFold(actual_name, name) {
				uuid := strings.Trim(s[2], " ")
				d.Set("uuid", uuid)

				serial := strings.Trim(s[1], " ")
				d.Set("adbserial", serial)
            }
        }
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return nil
}

func resourceGenymotionDelete(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)

	// Stop Genymotion Cloud Device
	cmd := exec.Command("gmtool", "admin", "--cloud", "stopdisposable", name )
	output, err := cmd.CombinedOutput()
	if err != nil {
    	return fmt.Errorf("Error: %s", output)
	} else {
    	fmt.Println(string(output))
	}
    return nil
}