build: 
	go build -o api cmd/main.go

run: build 
	./api

# Test
test:
	go test -v ./...

.PHONY: build run test