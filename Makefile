build:
	@go build -o bin/go-banking

run: build
	@./bin/go-banking

test:
	@go test -v ./...