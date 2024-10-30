.PHONY: dc run test lint

dc:
	docker-compose up  --remove-orphans --build

build:
	go build -race -o app cmd/main.go

run:
	go build -race -o app cmd/main.go && \
	HTTP_ADDR=:8080 \
	DEBUG_ERRORS=1 \
	DSN="postgres://postgres:@127.0.0.1:5432/bookshop?sslmode=disable" \
	MIGRATIONS_PATH="file://./internal/app/migrations" \
	./app

test:
	go test -race ./...

install-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.2

lint:
	golangci-lint run ./...

generate:
	go generate ./...