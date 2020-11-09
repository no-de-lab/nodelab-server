.PHONY: run vendor build

# load dotenv & export
include .env
export

# Install go modules dependencies
vendor:
	go mod vendor

run:
	@go run ./cmd/nodelabd/main.go

up:
	@docker-compose -f docker-compose.dev.yaml up

down:
	@docker-compose -f docker-compose.dev.yaml down -v

build:
	@CGO_ENABLED=0 GOOS=linux go build -o nodelabd ./cmd/...

# build for air
build-air:
	@go build -o ./tmp/app/engine cmd/nodelabd/main.go