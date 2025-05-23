APP_NAME = go-banking
ENTRY_POINT = cmd/api/main.go

build:
	@go build -o bin/${APP_NAME} ${ENTRY_POINT}

run: build
	@./bin/go-banking

test:
	@go test -v ./...