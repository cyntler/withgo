test:
	go test ./...

test-watch:
	gow test ./...

build:
	go build -o ./bin/withgo ./cmd/withgo
