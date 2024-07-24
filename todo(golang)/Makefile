.PHONY: build
build:
	go build -v ./cmd/apiserver && migrate -path migrations -database "postgres://postgres:qwe@localhost/db_golang?sslmode=disable" up
.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build