[![Concourse Terraform Provider](https://github.com/saurabh-mish/terraform-provider-concourse/actions/workflows/ci.yaml/badge.svg)](https://github.com/saurabh-mish/terraform-provider-concourse/actions/workflows/ci.yaml)

# Terraform Provider for Concourse Labs

This provider is currently under development and has not been published to the Terraform registry.

It uses an executable binary from the local terraform directory to perform operations on remote infrastructure.

## Getting Started

+ [Determine][1] the architecture of your operating system (`$GOARCH` column)

  Examples:

  + Mac with Intel processor - `darwin_amd64`
  + Mac with Apple processor - `darwin_arm64`
  + Windows (64-bit) with Intel processor - `windows_amd64`
  + Windows (64-bit) with AMD processor - `windows_arm64`
  + Linux (64-bit) with Intel processor - `linux_amd64`

+ Create a terraform plugin directory based on `$GOARCH`

  `~/.terraform.d/plugins/hashicorp.com/edu/concourse/0.3.1/<$GOARCH>`

  OR

  `%APPDATA%\terraform.d\plugins\hashicorp.com\edu\concourse\0.3.1\<$GOARCH>`

+ Clone this repository locally

  `git clone git@github.com:saurabh-mish/terraform-provider-concourse.git`

+ Navigate to project root

  `cd ./terraform-provider-concourse`

+ **Modify `OS_ARCH` and `BINARY` in the *Makefile* based on your operating system**

+ Run tests

  `make test`

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

+ The `data` block is used to retrieve existing resources.

+ The `resource` block is used to create, update, and delete attribute tags.


## Operations

+ Initialize

  `terraform init`

+ Plan

  `terraform plan`

+ Apply

  `terraform apply --auto-approve`

+ Destroy

  `terraform destroy --auto-approve`


---

[1]: https://go.dev/doc/install/source#environment
[2]: https://prod.concourselabs.io/
