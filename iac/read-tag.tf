# terraform {
#   required_providers {
#     concourse = {
#       version = "0.3.1"
#       source = "hashicorp.com/edu/concourse"
#     }
#   }
# }

# data "concourse_attribute_tag" "freddie_mac" {
#   id = 211012
# }

# data "concourse_attribute_tag" "first_tag" {
#   id = 60003
# }

# output "freddie" {
#   value = data.concourse_attribute_tag.freddie_mac
# }

# output "oldest" {
#   value = data.concourse_attribute_tag.first_tag
# }
