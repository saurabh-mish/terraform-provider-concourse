HOSTNAME=concourselabs.com
NAMESPACE=prod
TYPE=provider
NAME=concourse
BINARY=terraform-${TYPE}-${NAME}
VERSION=0.8
OS_ARCH=darwin_arm64

default: build

prep:
	go fmt ./...
	go mod tidy

test:
	go test ./concourse

build:
	go build -o ${BINARY}

install:
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
