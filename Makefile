#Challenge Makefile

start:
	go run src/main.go

check:
#TODO: include command to test the code and show the results

setup:
	docker build -t dci_pg pg_docker
	docker run -d --name dci_pg -p 5432:5432 dci_pg
	go get github.com/lib/pq
	go get github.com/gorilla/mux