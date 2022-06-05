terraform {
  required_providers {
    concourse = {
      version = "0.3.1"
      source = "hashicorp.com/edu/concourse"
    }
  }
}

//provider concourse {
//  username = "saurabh+113@concourselabs.com"
//  password = "S@r8hM"
//}

data "concourse_attribute_tag" "freddie_mac" {
  id = 211012
}

output "freddie" {
  value = data.concourse_attribute_tag.freddie_mac
}
