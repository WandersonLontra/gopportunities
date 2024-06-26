.PHONY: default run build test docs clean

# variables

APP_NAME=gopportunities

# tasks

default: run


run:
	@go run ./cmd/gopportunities/main.go
build:
	@go build -o $(APP_NAME) ./cmd/gopportunities/main.go
test:
	@go test ./ ...
docs:
	@cd ./cmd/gopportunities
	@swag init -g ./cmd/gopportunities/main.go -o docs
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs

	