terraform {
  required_providers {
    concourse = {
      version = "0.8"
      source = "concourselabs.com/prod/concourse"
    }
  }
}

# create, update, and delete operations
resource "concourse_attribute_tag" "demo" {
	name = "test name"
	description = "test description"
}

# read operation
data "concourse_attribute_tag" "custom_tag" {
  id = concourse_attribute_tag.demo.id
}

# print to stdout
output "attr_tag" {
  value = data.concourse_attribute_tag.custom_tag
}
