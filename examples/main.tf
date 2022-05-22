terraform {
  required_providers {
    hashicups = {
      version = "0.3.1"
      source  = "hashicorp.com/edu/hashicups"
    }
  }
}

provider "hashicups" {}

//provider "hashicups" {
//  username = "education"
//  password = "test123"
//}

data "hashicups_order" "order" {
  id = 1
}

output "order" {
  value = data.hashicups_order.order
}
