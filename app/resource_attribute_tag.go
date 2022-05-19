package app

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/saurabh-mish/terraform-provider-concourse/client"
)

func resourceAttributeTag() *schema.Resource {
	return &schema.Resource {
		Create: resourceItemCreate,
		Read: resourceItemRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceAttributeTagCreate(data *schema.ResourceData, meta interface{}) error {

	apiClient := meta.(*client.Client)
	item := client.Item{}

	if val, ok := data.GetOk(key: "name"); ok {
		item.Name = val.(string)
	}

	if val, ok := data.GetOk(key: "description"); ok {
		item.Description = val.(string)
	}

	if val, ok := data.GetOk(key: "tags"); ok {
		tfTags := val.(*schema.Set).List()
		tags := make([]string, len(tfTags))
		for i, tfTag := range tfTags {
			tags[i] = tfTag.(string)
		}
	}

	err := apiClient.NewItem(&item)
	if err != nil {
		return err
	}

	data.SetId(item.Name)

	return nil
}

func resourceAttributeTagRead(data *schema.ResourceData, meta interface{}) error {

	apiClient := meta.(*client.Client)
	itemId := data.Id()

	item, err := apiClient.getItem(itemId)
	if err != nil {
		/*if string.Contains(err.Error(), "not found") {
			data.setId("")
		} else {
			return fmt.Errorf("error finding item with ID", )
		}*/
		return err
	}

	data.Set("name": Item.Name)
	data.Set("description", Item.Description)
	data.Set("tags": Item.Tags)

	return nil
}
