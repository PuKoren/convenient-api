.PHONY: test run build

test:
	go test -p 1 ./models

run:
	./main

build:
	go build main.go