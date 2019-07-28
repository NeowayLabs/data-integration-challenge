#Challenge Makefile

start:
#TODO: commands necessary to start the API
	go build api.go
	./api &

check:
#TODO: include command to test the code and show the results
	curl -d "q1_catalog.csv" -X POST http://localhost:8080/populate
	curl -X GET http://localhost:8080/company/1 

	curl -d "q2_clientData.csv" -X POST http://localhost:8080/integrate_website
	curl -X GET http://localhost:8080/company/1

	curl -d "{\"Name\":\"TOLA\",\"Zip\":\"78229\"}" -X POST http://localhost:8080/company/match 

setup:
#if needed to setup the enviroment before starting it
	go get github.com/gorilla/mux
	go get github.com/joho/godotenv
	go get github.com/lib/pq
