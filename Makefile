.PHONY:build
echo:
	@echo "hello world"

init:
	go mod tidy
	go mod download
	go mod vendor

build:
	go build -v ./...

run:
	go run ./cmd/abc/main.go

build-abc:
	DOCKER_BUILDKIT=0 docker build -f build/abc/Dockerfile . -t example/abc

run-abc:
	docker run -it --rm -p 8081:8080 example/abc

deploy:
	kubectl apply -f deployments/kubernetes/abc

test:
	go test -v ./...
