.PHONY:build
echo:
	@echo "hello world"

init:
	go env -w GOPROXY=https://goproxy.cn,direct
	go mod tidy
	go mod download
	go mod vendor

build:
	go build -v ./cmd/example

run:
	go run ./cmd/example/main.go

build-example:
	DOCKER_BUILDKIT=0 docker build -f build/example/Dockerfile . -t example/example

run-example:
	docker run -it --rm -p 8081:8080 example/example

deploy:
	kubectl apply -f deployments/kubernetes/example

test:
	go test -v ./...
