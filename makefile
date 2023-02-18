build: 
	go build -o api cmd/main.go

run: build 
	./api