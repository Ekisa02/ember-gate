run:
	go run ./cmd/server

build:
	go build -o bin/ember-gate ./cmd/server

test:
	go test ./... -v
