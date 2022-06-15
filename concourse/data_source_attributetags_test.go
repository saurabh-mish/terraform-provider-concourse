package concourse

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAcceptanceAttributeTagDataSource(t *testing.T) {
	rInt := rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAttrTagDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAcceptanceAttributeTagDataSourceConfig(rInt),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"concourse_attribute_tag.demo_tag", "id",
						"data.concourse_attribute_tag.new_tag", "id",
					),
				),
			},
		},
	})
}

func testAcceptanceAttributeTagDataSourceConfig(rInt int) string {
	return fmt.Sprintf(`
	terraform {
      required_providers {
        concourse = {
          version = "0.3.1"
          source = "hashicorp.com/edu/concourse"
        }
      }
    }

    resource "concourse_attribute_tag" " demo_tag" {
      name = "Concourse custom provider"
      description = "Concourse custom provider description"
    }

    data "concourse_attribute_tag" "new_tag" {
      id = concourse_attribute_tag.demo_tag.id
    }
	`)
}
