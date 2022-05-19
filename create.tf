provider "hashicups" {
  username = var.username
  password = var.password
}

variable "username" {
  type        = string
  description = "HashiCups login username"
  sensitive   = false
}

variable "password" {
  type        = string
  description = "HashiCups login password"
  sensitive   = true
}


resource "hashicups_order" "edu" {
  items {
    coffee {
      id = 3
    }
    quantity = 3
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
