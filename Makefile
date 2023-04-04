build:
	go build cmd/main.go

build-race:
	go build -race cmd/main.go

run-local:
	go build cmd/main.go && ./main

init:
	go install github.com/swaggo/swag/cmd/swag@latest && swag init

test-local:
	go test -v test/* && go test -race test/*

run-local-race:
	go build -race cmd/main.go && ./main
