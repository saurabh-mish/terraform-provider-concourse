# Hashicups Provider

Terraform plugins follow naming convention of type `terraform-<TYPE>-<NAME>`

## Run API locally

+ Start Docker Desktop GUI

+ Run API

  `cd docker_compose && docker compose up`

+ Verify that API is running (**in a new terminal**)

  `curl localhost:19090/health`

## Provider

+ Define directory as module root

  `go mod init github.com/saurabh-mish/terraform-provider-hashicups`

+ Install dependencies

  `go mod tidy`

+ Build provider using Makefile

  `make build`

+ Execute binary

  `./terraform-provider-hashicups`

+ Move to local plugin directory

  `mv terraform-provider-hashicups ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.3.1/darwin_arm64`

## Terraform

+ Initialize

  `terraform init`

+ Apply read

  `terraform apply --auto-approve`
