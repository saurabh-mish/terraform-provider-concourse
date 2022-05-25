# Hashicups Provider

Performs read operation on a locally running API.

## Run API locally

+ Start Docker Desktop GUI

+ Run API

  `cd docker_compose && docker compose up`

+ Verify that API is running (**in a new terminal**)

  `curl localhost:19090/health`

## REST Endpoints

+ User Auth

  + Sign-up

    `curl -X POST localhost:19090/signup -d '{"username":"education", "password":"test123"}'`

  + Sign-in

    `curl -X POST localhost:19090/signin -d '{"username":"education", "password":"test123"}'`

  + Set environment variables

    ```
    export HASHICUPS_USERNAME=education
    export HASHICUPS_PASSWORD=test123
    export HASHICUPS_TOKEN="ey ..."
    ```

+ Read

  + All coffee

    `curl localhost:19090/coffees`

  + Order ID

    `curl -X GET -H "Authorization: ${HASHICUPS_TOKEN}" localhost:19090/orders/<ID>`

+ Create

  `curl -X POST -H "Authorization: ${HASHICUPS_TOKEN}" localhost:19090/orders -d '[{"coffee": { "id":1 }, "quantity":2}, {"coffee": { "id":3 }, "quantity":2}]'`

+ Update

  `curl -X PUT -H "Authorization: ${HASHICUPS_TOKEN}" localhost:19090/orders/<ID> -d  '[{"coffee": { "id":1 }, "quantity":1}, {"coffee": { "id":3 }, "quantity":3}]'`

+ Delete

  `curl -X DELETE -H "Authorization: ${HASHICUPS_TOKEN}" localhost:19090/orders/<ID>`

## Data Structure

The *order* service requires coffee ID and quantity.

It has the following structure:

```
{
  "id": 7,                                 --> order ID
  "items": [
    {
      "coffee": {
        "id": 1,                           --> coffee ID
        "name": "Packer Spiced Latte",
        "teaser": "Packed with goodness to spice up your images",
        "description": "",
        "price": 350,
        "image": "/packer.png",
        "ingredients": null
      },
      "quantity": 1
    },
    {
      "coffee": {
        "id": 3,
        "name": "Nomadicano",
        "teaser": "Drink one today and you will want to schedule another",
        "description": "",
        "price": 150,
        "image": "/nomad.png",
        "ingredients": null
      },
      "quantity": 3
    }
  ]
}
```

## Provider

+ Create directory to store binary

  `mkdir ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.3.1/darwin_arm64`

+ Define directory as module root

  `go mod init github.com/saurabh-mish/terraform-provider-concourse`

+ Install dependencies

  `go mod tidy`

+ Build provider and install locally using Makefile

  ```
  make build
  make install
  ```

+ Execute binary

  `./terraform-provider-concourse`

## Terraform

+ Change to `examples/` directory (from project root)

  `cd examples/`

+ Initialize

  `terraform init`

+ Apply read

  `terraform apply --auto-approve`

## Miscellaneous

#### Docker Cleanup

+ Delete all images

  `docker rmi -f $(docker images -aq)`

+ Delete all containers including its volumes

  `docker rm -vf $(docker ps -aq)`

+ Remove all unused containers, volumes, networks and images

  `docker system prune -a --volumes`
