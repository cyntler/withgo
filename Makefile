run:
	go run ./cmd/withgo $(ARGS)

test:
	go test ./...

test-watch:
	@>/dev/null which gow || (echo "Missing gow, install with: go install github.com/mitranim/gow" && false)
	gow test ./...

build:
	go build -o ./bin/withgo ./cmd/withgo
