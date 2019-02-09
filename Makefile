#Challenge Makefile

start:
	go run $(ls *.go | grep -v 'test')

check:
	go test ./... -cover

setup:
		/bin/bash setup.sh
