build:
	@go build -o bin/main main.go

run: build
	@bin/main

test: 
	@go test -v ./...

itest:
	@py integration/main.py