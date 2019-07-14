#Challenge Makefile

start:
	go run main.go

check:
	go test

setup :
	go get ./...
