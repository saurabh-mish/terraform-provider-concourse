HOSTNAME=hashicorp.com
NAMESPACE=edu
TYPE=provider
NAME=concourse
BINARY=terraform-${TYPE}-${NAME}
VERSION=0.3.1
OS_ARCH=darwin_arm64

default: build

format:
	go fmt ./...

build:
	go build -o ${BINARY}

install:
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
