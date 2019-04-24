package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema:         providerSchema(),
		ResourcesMap:   providerResources(),
		DataSourcesMap: providerDataSources(),
		ConfigureFunc:  providerConfigure,
	}
}

// List of supported configuration fields for the provider.
// More info in https://github.com/hashicorp/terraform/blob/v0.6.6/helper/schema/schema.go#L29-L142
func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"host": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Host address of the SmartOS global zone.",
		},
		"user": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "User to authenticate with.",
		},
	}
}

func providerResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"smartos_machine": resourceMachine(),
	}
}

func providerDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"smartos_image": datasourceImage(),
	}
}

// This is the function used to fetch the configuration params given
// to our provider which we will use to initialise a client that
// interacts with the Project Fifo API.
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	client := SmartOSClient{
		host: d.Get("host").(string),
		user: d.Get("user").(string),
	}

	return &client, nil
}
