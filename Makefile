build:
	go build -o bin/aoc

run: build
	./bin/aoc

test: 
	go test -v ./...
	