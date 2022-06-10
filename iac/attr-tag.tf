terraform {
  required_providers {
    concourse = {
      version = "0.3.1"
      source = "hashicorp.com/edu/concourse"
    }
  }
}

# create and delete operations
resource "concourse_attribute_tag" "tagtest" {
	name = "saurabh test change"
	description = "saurabh terraform provider test description"
}
