terraform {
  required_providers {
    concourse = {
      version = "0.3.1"
      source = "hashicorp.com/edu/concourse"
    }
  }
}

resource "concourse_attribute_tag" "tagtest" {
	name = "saurabh test name"
	description = "saurabh test description"
}
