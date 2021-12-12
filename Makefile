.PHONY: install test cover run.dev build

install:
	go mod download

test:
	mkdir -p ./coverage && \
		go test -v -coverprofile=./coverage/coverage.out -covermode=atomic ./...

cover: test
	go tool cover -func=./coverage/coverage.out &&\
		go tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html

run.dev:
	go run ./app/main.go

build:
	go build app/main