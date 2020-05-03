GOPATH:=$(shell go env GOPATH)
tag := v1.0.7
registryUrl := registry.cn-shanghai.aliyuncs.com
registryName := chenwentao
projectName := seat
imageName:= ${registryUrl}/${registryName}/${projectName}:${tag}
.PHONY: build
echo:
	echo ${imageName}
build_xprocess:
	go build -o bin/xprocess-server
build_xprocess_linux:
	GOOS=linux GOARCH=amd64 go build -o bin/xprocess-server
.PHONY: docker
docker:
	docker build . -t ${imageName}
push:
	docker push ${imageName}
removeIamge:
	docker rmi ${imageName}
all:
	GOOS=linux GOARCH=amd64 go build -o bin/process-server
	docker build . -t ${imageName}
	docker push ${imageName}
	docker rmi ${imageName}
