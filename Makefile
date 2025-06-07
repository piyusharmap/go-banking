APP_NAME = go-banking
ENTRY_POINT = cmd/api/main.go
MIGRATION_DIR = db/migrations

build:
	@go build -o bin/${APP_NAME} ${ENTRY_POINT}

run: build
	@./bin/go-banking

test:
	@echo "Running tests"
	@go test -v ./...

create-migrate:
	@echo "Creating new migration: ${NAME}"
	@migrate create -ext sql -dir ${MIGRATION_DIR} ${NAME}

migrate-up:
	@migrate -database  "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path ${MIGRATION_DIR} -verbose up

migrate-down:
	@migrate -database  "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path ${MIGRATION_DIR} -verbose down 1

force-version:
	@migrate -database  "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path ${MIGRATION_DIR} force ${VERSION}