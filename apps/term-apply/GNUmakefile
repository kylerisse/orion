build:	*.go
	go build -o term-apply

term-apply: build

deps:
	go mod verify
	go mod tidy

test:
	go clean -testcache
	go test --race ./...
