GOPATH:=$(shell go env GOPATH)

.PHONY: build
build_process:
	go build -o bin/process-server

build_process_linux:
	GOOS=linux GOARCH=amd64 go build -o bin/process-server
.PHONY: docker
docker:
	docker build . -t registry.cn-shanghai.aliyuncs.com/chenwentao/xprocess-backend:0.0.9

