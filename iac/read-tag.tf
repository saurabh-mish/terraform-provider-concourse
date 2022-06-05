terraform {
  required_providers {
    concourse = {
      version = "0.3.1"
      source = "hashicorp.com/edu/concourse"
    }
  }
}

data "attribute_tag" "freddie_mac" {
  id = 211012
}

output "freddie" {
  value = data.attribute_tag.freddie_mac
}
