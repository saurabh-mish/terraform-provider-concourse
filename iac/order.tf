terraform {
  required_providers {
    hashicups = {
      version = "0.3.1"
      source = "hashicorp.com/edu/hashicups"
    }
  }
}

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

resource "hashicups_order" "edu" {
  items {
    coffee {
      id = 3
    }
    quantity = 1
  }
  items {
    coffee {
      id = 2
    }
    quantity = 1
  }
}

output "edu_order" {
  value = hashicups_order.edu
}


data "hashicups_order" "first" {
  id = 1
}

output "first_order" {
  value = data.hashicups_order.first
}
