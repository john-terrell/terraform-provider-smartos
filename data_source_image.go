package main

import (
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

	image, err := client.GetImage(name, version)
	if err != nil {
		log.Printf("Failed to retrieve image with name: %s, version: %s.  Error: %s", name, version, err)
		return err
	}

	d.SetId(image.ID.String())

	return nil
}
