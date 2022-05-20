TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=hashicorp.com
NAMESPACE=edu
TYPE=provider
NAME=hashicups
BINARY=terraform-${TYPE}-${NAME}
VERSION=0.3.1
OS_ARCH=darwin_arm64

default: build

build:
	go build -o ${BINARY}
