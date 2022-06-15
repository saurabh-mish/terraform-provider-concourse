package concourse

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/saurabh-mish/terraform-provider-concourse/client"
)

func dataSourceAttributeTag() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAttributeTagRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_by": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"updated_by": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"institution_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAttributeTagRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	var diags diag.Diagnostics

	tagID := strconv.Itoa(d.Get("id").(int))

	resp, err := c.GetAttributeTag(tagID)
	if err != nil {
		return diag.FromErr(err)
	}

	//d.Set("id", resp.ID)
	d.Set("version", resp.Version)
	d.Set("created", resp.Created)
	d.Set("updated", resp.Updated)
	d.Set("created_by", resp.CreatedBy)
	d.Set("updated_by", resp.UpdatedBy)
	d.Set("institution_id", resp.InstitutionId)
	d.Set("name", resp.Name)
	d.Set("description", resp.Description)

	// set response body
	d.SetId(tagID)

	return diags
}
