#!make
start:
	go run ./cmd/usher/main.go

build:
	 go build -o ./app/usher ./cmd/usher

lint:
	golangci-lint run --timeout 10m

test: test_unit test_e2e

test_unit:
	go test -v -cover -race -short  -coverprofile=coverage.out ./...

test_e2e:
	go test -v -run E2E ./test/...

benchmark:
	go test -benchmem -bench=. ./...	

coverprofile:
	go tool cover -html=coverage.out

openapi:
	bash ./scripts/openapi_gen.sh

entgo:
	go generate ./internal/infrastructure/store/ent