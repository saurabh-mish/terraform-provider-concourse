# Hashicups Provider

Performs read operation on a locally running API.

## Run API locally

+ Start Docker Desktop GUI

+ Run API

  `cd docker_compose && docker compose up`

+ Verify that API is running (**in a new terminal**)

  `curl localhost:19090/health`

## Data

+ Signin

  `curl -X POST localhost:19090/signin -d '{"username":"education", "password":"test123"}'`

+ Save token

  `export HASHICUPS_TOKEN=ey.....`

+ Add data

  `curl -X POST -H "Authorization: ${HASHICUPS_TOKEN}" localhost:19090/orders -d '[{"coffee": { "id":1 }, "quantity":4}, {"coffee": { "id":3 }, "quantity":3}]'`

+ Verify data

  `curl -X GET -H "Authorization: ${HASHICUPS_TOKEN}" localhost:19090/orders/1`

## Provider

  + Define directory as module root

    `go mod init github.com/saurabh-mish/terraform-provider-hashicups`

  + Install dependencies

    `go mod tidy`

  + Build provider and install locally using Makefile

    ```
    make build
    make install
    ```

  + Execute binary

    `./terraform-provider-hashicups`

## Terraform

+ Change to `examples/` directory (from project root)

  `cd examples/`

+ Initialize

  `terraform init`

+ Apply read

  `terraform apply --auto-approve`
