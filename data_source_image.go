package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func datasourceImage() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Read:          datasourceImageReadRunc,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func datasourceImageReadRunc(d *schema.ResourceData, m interface{}) error {
	client := m.(*SmartOSClient)

	name := d.Get("name").(string)
	version := d.Get("version").(string)

	var image *Image
	var err error

	image, err = client.GetLocalImage(name, version)
	if err != nil {
		return err
	}

	if image == nil {
		image, err = client.FindRemoteImage(name, version)
		if err == nil && image == nil {
			return fmt.Errorf("Image not found")
		}
	}

	if err != nil {
		log.Printf("Failed to retrieve image with name: %s, version: %s.  Error: %s", name, version, err)
		return err
	}

	d.SetId(image.ID.String())

	return nil
}
