build:
	@go build -o bin/standardApi cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/standardApi