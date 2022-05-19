package app

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Resources required to maintain connection with the API endpoint
func Provider() terraform.ResourceProvider {
	return &schema.Provider {
		Schema: map[string]*schema.Schema{
			"token": {
				Type:     schema.TypeString,
				Required: true
				Description: "JSON web token (JWT)"
				DefaultFunc: schema.EnvDefaultFunc("CONCOURSE_TOKEN", "")
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"app_item": resourceItem(),
		},
		ConfigureFunc: configureProvider,
	}
}

// Resources required to create connection with the API endpoint
func configureProvider(d *schema.ResourceData) (interface{}, error) {
	token := d.Get(key: "token").(string)
	return client.NewClient(token), nil
}
