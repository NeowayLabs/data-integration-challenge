#Challenge Makefile

start-server:
	go run ./cmd/server/main.go

check:
	go test ./... -cover

setup:
		/bin/bash scripts/setup.sh
