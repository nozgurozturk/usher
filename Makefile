#!make
start:
	go run ./cmd/usher/main.go

build:
	 go build -o ./app/usher ./cmd/usher

lint:
	golangci-lint run --timeout 10m

test: test_unit test_integration

test_unit:
	go test -v -cover -race -short  -coverprofile=coverage.out ./...

test_integration:
	go test -v -run Integration ./...

coverprofile:
	go tool cover -html=coverage.out