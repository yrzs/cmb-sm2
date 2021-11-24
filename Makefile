GOPATH:=$(shell go env GOPATH)
API_PROTO_FILES=$(shell find api -name *.proto)
OUTPUT_PYTHON =$(shell rm -fr ./scripts/output/python/*)./scripts/output/python
OUTPUT_PHP =$(shell rm -fr ./scripts/output/php/*)./scripts/output/php
VERSION = xc/achi-rank:v1

.PHONY: test
test:
	echo "test"

.PHONY: api
api:
	protoc --proto_path=. \
	       --proto_path=./api \
           --proto_path=/usr/local/include/google \
 	       --go_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative:. \
	       $(API_PROTO_FILES)

.PHONY: py
py:
	python  -m  grpc_tools.protoc  --python_out=$(OUTPUT_PYTHON) --grpc_python_out=$(OUTPUT_PYTHON)  -I=. $(API_PROTO_FILES)

.PHONY: php
php:
	protoc --proto_path=./api --php_out=$(OUTPUT_PHP) --grpc_out=$(OUTPUT_PHP) -I=. $(API_PROTO_FILES)

.PHONY: build
build:
	mkdir -p bin/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/ main.go

.PHONY: run
run:
	./bin/main

.PHONY: bar
bar:
	make build && make run

.PHONY: docker
docker:
	make build && docker build -t $(VERSION) .