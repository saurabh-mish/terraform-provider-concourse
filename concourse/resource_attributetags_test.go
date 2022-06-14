package concourse

import (
	"fmt"
	"testing"

	"github.com/saurabh-mish/terraform-provider-concourse/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccAttributeTag_basic(t *testing.T) {
	resourceName := "concourse_attribute_tag.demo_tag"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAttrTagDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAcceptanceAttributeTag_config(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test name"),
					resource.TestCheckResourceAttr(resourceName, "description", "test description"),
				),
			},
		},
	})
}

func testAccCheckAttrTagDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*client.Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "concourse_attribute_tag" {
			continue
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No instance ID is set")
		}
		_, err := conn.GetAttributeTag(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("Attribute Tag %s still exists", rs.Primary.ID)
		}
	}
	return nil
}

func testAcceptanceAttributeTag_config() string {
	return fmt.Sprintf(`

	terraform {
      required_providers {
        concourse = {
          version = "0.3.1"
          source = "hashicorp.com/edu/concourse"
        }
      }
    }

	resource "concourse_attribute_tag" "demo_tag" {
	  name = "test name"
	  description = "test description"
	}
	`)
}
