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
			"createdBy": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"updatedBy": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"institutionId": &schema.Schema{
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

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	tagID := strconv.Itoa(d.Get("id").(int))

	_, err := c.GetAttributeTag(tagID)
	if err != nil {
		return diag.FromErr(err)
	}

	// set response body
	d.SetId(tagID)

	return diags
}

/*
func dataSourceAttributeTagRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  client := &http.Client{Timeout: 10 * time.Second}

  // Warning or errors can be collected in a slice type
  var diags diag.Diagnostics

  req, err := http.NewRequest("GET", fmt.Sprintf("%s/coffees", "http://localhost:19090"), nil)
  if err != nil {
    return diag.FromErr(err)
  }

  r, err := client.Do(req)
  if err != nil {
    return diag.FromErr(err)
  }
  defer r.Body.Close()

  coffees := make([]map[string]interface{}, 0)
  err = json.NewDecoder(r.Body).Decode(&coffees)
  if err != nil {
    return diag.FromErr(err)
  }

  if err := d.Set("coffees", coffees); err != nil {
    return diag.FromErr(err)
  }

  // always run
  d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

  return diags
}
*/
