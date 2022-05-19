# HashiCups

Terraform provider for a fictional coffee-shop application


## Initialize HashiCups locally

+ Start local instance on port `:19090`

  `cd docker_compose && docker compose up`

+ In a new terminal (tab) check the health endpoint

  `curl localhost:19090/health`

  **NOTE**: The below message will be logged in the initial terminal:

    `docker_compose-api-1  | 2022-05-19T17:52:48.812Z [INFO]  Starting service: bind=0.0.0.0:9090 metrics=localhost:9102`


## Provider

The Hashicups provider is not published to the Terraform registry because its used only for local developoment and testing.

To be able to use this provider, it has to be present locally in:
`~/.terraform.d/plugins/${host_name}/${namespace}/${type}/${version}/${target}`.

In this case, the path will translate to:<br>
`~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.3.1/darwin_arm64`

+ Create directory

  `mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.3.1/darwin_arm64 \
  && cd ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.3.1/darwin_arm64`

+ Download and unzip provider, make it executable

  ```
  curl -LO https://github.com/hashicorp/terraform-provider-hashicups/releases/download/v0.3.1/terraform-provider-hashicups_0.3.1_darwin_amd64.zip

  unzip terraform-provider-hashicups_0.3.1_darwin_arm64.zip

  rm README terraform-provider-hashicups_0.3.1_darwin_arm64.zip

  chmod +x terraform-provider-hashicups_v0.3.1
  ```


## Application

+ Authenticate with the `/signup` endpoint with credentials `education` and `test123` respectively

  `curl -X POST localhost:19090/signup -d '{"username":"education", "password":"test123"}'`

+ Set environment variable `HASHICUPS_TOKEN` to the token value received as response

  `export HASHICUPS_TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTMwNzAzNDksInVzZXJfaWQiOjEsInVzZXJuYW1lIjoiZWR1Y2F0aW9uIn0.EIDanBdUg_nkSybM3tO4Z9jkedRCpWu76ZqPrRNN7t0`

  **NOTE**: The below message will be logged in the initial terminal:

    `docker_compose-api-1  | 2022-05-19T18:12:29.744Z [INFO]  Handle User | signup`


## Terraform

Initialize Terraform:<br>
`terraform init`

Since HashiCups is a third-party provider, the `hostname` and `namespace` values in the source string are arbitrary.
