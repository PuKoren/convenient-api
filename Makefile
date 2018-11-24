.PHONY: test run build

build:
	go build main.go

test:
	go test -p 1 ./models

deps:
	dep ensure

run:
	./main
