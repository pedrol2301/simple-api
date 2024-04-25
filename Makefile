build:
	@go build -o bin/simple-api cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/simple-api