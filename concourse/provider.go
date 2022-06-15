package concourse

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/saurabh-mish/terraform-provider-concourse/client"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CONCOURSE_USERNAME", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("CONCOURSE_PASSWORD", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"concourse_attribute_tag": resourceAttributeTag(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"concourse_attribute_tag": dataSourceAttributeTag(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	var diags diag.Diagnostics // slice for warnings and errors

	if (username != "") && (password != "") {
		c, err := client.NewClient(nil, &username, &password)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Credentials not present",
				Detail:   "Environment variables for username and password dont exist",
			})
			return nil, diags
		}
		return c, diags
	}

	c, err := client.NewClient(nil, nil, nil)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unauthorized access",
			Detail:   "Unable to authenticate user with Concourse",
		})
		return nil, diags
	}
	return c, diags
}
