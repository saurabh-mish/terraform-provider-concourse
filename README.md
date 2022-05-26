# concourse

Terraform provider for Concourse Labs

## Data Structures

#### Authentication

Authenticate to get a user token

```
curl --request POST 'https://auth.prod.concourselabs.io/api/v1/oauth/token' \
  --header 'Accept: application/json' \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --data-urlencode 'username=user+113@concourselabs.com' \
  --data-urlencode 'password=decentPassword' \
  --data-urlencode 'grant_type=password' \
  --data-urlencode 'scope=INSTITUTION POLICY MODEL IDENTITY RUNTIME_DATA' \
  | jq
```
