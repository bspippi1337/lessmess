.PHONY: tidy run build

tidy:
	go mod tidy

run:
	go run ./cmd/restless

build:
	mkdir -p bin
	go build -o bin/restless ./cmd/restless
