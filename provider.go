package main

import (
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
        return &schema.Provider{
			Schema: map[string]*schema.Schema{},

			DataSourcesMap: map[string]*schema.Resource{
				"chucknorris": dataSource(),
			},			
			
			ResourcesMap: map[string]*schema.Resource{},
        }
}
