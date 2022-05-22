# Hashicups Provider

Performs read operation on a locally running API.

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

+ Build provider and install locally using Makefile

  `make build`

+ Execute binary

  `./terraform-provider-hashicups`

## Terraform

+ Change to `examples/` directory (from project root)

  `cd examples/`

+ Initialize

  `terraform init`

+ Apply read

  `terraform apply --auto-approve`
