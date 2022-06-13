terraform {
  required_providers {
    concourse = {
      version = "0.3.1"
      source = "hashicorp.com/edu/concourse"
    }
  }
}

# create, update, and delete operations
resource "concourse_attribute_tag" "demo" {
	name = "saurabh test name"
	description = "saurabh test description"
}

# read operation
data "concourse_attribute_tag" "custom_tag" {
  id = concourse_attribute_tag.demo.id
}

# print to stdout
output "attr_tag" {
  value = data.concourse_attribute_tag.custom_tag
}
