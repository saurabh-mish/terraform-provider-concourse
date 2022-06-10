# Terraform Provider for Concourse Labs

This provider is currently under development and has not been published to the Terraform registry.

It uses an executable binary from the local terraform directory to perform operations on remote infrastructure.

## Getting Started

+ [Determine][1] the architecture of your operating system (`$GOARCH` column)

  Examples:

  + Mac with Intel processor would be `darwin_amd64`
  + Mac with Apple processor would be `darwin_arm64`

+ Create a terraform plugin directory based on `$GOARCH`

  `mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/concourse/0.3.1/<$GOARCH>`

+ Modify `OS_ARCH` in the Makefile


+ Build the binary

  `make build`

+ Move the binary to the local terraform plugin directory

  `make install`

## Infrastructure as Code

This release supports read operation. It is being expanded to support create, update, and delete operations.

+ Set environment variables corresponding to your [Concourse login][2]

  ```
  export CONCOURSE_USERNAME="user+113@concourselabs.com"
  export CONCOURSE_PASSWORD="decent_Password"
  ```

  **Note that this could be set in the provider block (but is not a security best practice)**

+ Navigarte to the `iac/` directory from the project root

  `cd ./iac`

+ Remove all terraform state files, cache files, etc. as they could be associated with a different build version of the provider

  `rm -rf .terraform .terraform.lock.hcl terraform.tfstate terraform.tfstate.backup`

+ The `data` block in `read-tag.tf` is used to retrieve existing resources. Two such IDs are predefined for demo

+ The `resource` block in the `attr-tag.tf` is used to create, update, and delete attribute tags.


## Terraform

+ Initialize Terraform

  `terraform init`

+ Plan

  `terraform plan`

+ Apply

  `terraform apply --auto-approve`

+ Destroy (update local TF files)

  `terraform destroy --auto-approve`


---

[1]: https://go.dev/doc/install/source#environment
[2]: https://prod.concourselabs.io/
