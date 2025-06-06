APP_NAME = go-banking
ENTRY_POINT = cmd/api/main.go
MIGRATION_DIR = db/migrations

build:
	@go build -o bin/${APP_NAME} ${ENTRY_POINT}

run: build
	@./bin/go-banking

test:
	@go test -v ./...

migrate-up:
	@migrate -database  "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path ${MIGRATION_DIR} -verbose up

migrate-down:
	@migrate -database  "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path ${MIGRATION_DIR} -verbose down 1

force-version:
	@migrate -database  "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path ${MIGRATION_DIR} force ${VERSION}