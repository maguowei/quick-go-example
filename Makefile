.PHONY:build
echo:
	@echo "hello world"

init:
	go env -w GOPROXY=https://goproxy.cn,direct
	go mod tidy
	go mod download
	go mod vendor

build:
	go build -v ./...

run:
	go run ./cmd/hello/main.go

build-hello:
	DOCKER_BUILDKIT=0 docker build -f build/hello/Dockerfile . -t example/hello

run-hello:
	docker run -it --rm -p 8081:8080 example/hello

deploy:
	kubectl apply -f deployments/kubernetes/hello

test:
	go test -v ./...
