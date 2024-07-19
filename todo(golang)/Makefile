.PHONY: build
build:
	go build -v ./cmd/apiserver && migrate -path migrations -database "postgres://postgres:qwe@localhost/db_golang?sslmode=disable" up
.DEFAULT_GOAL := build
