run:
	go run ./cmd/flucli

test:
	go test ./...

build:
	go build -o bin/flucli ./cmd/flucli

api:
	go run ./cmd/server