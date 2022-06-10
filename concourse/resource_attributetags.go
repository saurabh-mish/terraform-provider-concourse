package concourse

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/saurabh-mish/terraform-provider-concourse/client"
)

func resourceAttributeTag() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAttributeTagCreate,
		ReadContext:   resourceAttributeTagRead,
		UpdateContext: resourceAttributeTagUpdate,
    	DeleteContext: resourceAttributeTagDelete,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}


func resourceAttributeTagCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	var diags diag.Diagnostics

	attrTagData := client.AttrTagReq{
		Name: d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	resp, err := c.CreateAttributeTag(attrTagData)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(resp.ID))

	resourceAttributeTagRead(ctx, d, m)

	return diags
}

func resourceAttributeTagRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	c := m.(*client.Client)
	var diags diag.Diagnostics

	tagID := d.Id()

	resp, err := c.GetAttributeTag(tagID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("id", resp.ID)
	d.Set("version", resp.Version)
	d.Set("created", resp.Created)
	d.Set("updated", resp.Updated)
	d.Set("created_by", resp.CreatedBy)
	d.Set("updated_by", resp.UpdatedBy)
	d.Set("institution_id", resp.InstitutionId)
	d.Set("name", resp.Name)
	d.Set("description", resp.Description)

	return diags
}


func resourceAttributeTagUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  return resourceAttributeTagRead(ctx, d, m)
}

func resourceAttributeTagDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  var diags diag.Diagnostics

  return diags
}
