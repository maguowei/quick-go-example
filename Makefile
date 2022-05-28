.PHONY:build
echo:
	@echo "hello world"

init:
	go env -w GOPROXY=https://goproxy.cn,direct
	go mod download
	go run -mod=mod entgo.io/ent/cmd/ent generate ./internal/example/domain/repository/ent/schema
	go mod tidy

build:
	go build -v ./cmd/example

protoc:
	protoc -I . --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/proto/*.proto

run:
	go run ./cmd/example/main.go

build-example:
	DOCKER_BUILDKIT=0 docker build -f build/example/Dockerfile . -t example/example

run-example:
	docker run -it --rm -p 8081:8080 example/example

deploy:
	kubectl apply -f deployments/kubernetes/example

del-deploy:
	kubectl delete -f deployments/kubernetes/example

test:
	go test -v ./...
