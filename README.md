# Hashicups Provider

Terraform plugins follow naming convention of type `terraform-<TYPE>-<NAME>`

## Build provider

+ Define directory as odule root

  `go mod init terraform-provider-hashicups`

+ Install provider dependencies

  `go mod tidy`

+ Build provider using Makefile

  `make build`

+ Execute binary

  `./terraform-provider-hashicups`
