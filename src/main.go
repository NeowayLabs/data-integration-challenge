package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

var chanParceCsv = make(chan Companies, 4)

var chanDone = make(chan bool)

const (
	DB_USER     = "dci"
	DB_PASSWORD = "dci"
	DB_NAME     = "dci"
	DB_PORT     = "5432"
)

func main() {

	setupRoutes()

}

func processFile(csvFile *os.File) {

	fmt.Println("process file ", csvFile.Name())

	dbinfo := fmt.Sprintf("port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		panic(err)
	}

	db.Ping()

	go consumeChainParcerCsv(db)

	go parseCsvFile(csvFile)

	<-chanDone

}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("", "upload.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	// read all of the contents of our uploaded file into a
	// byte array
	fmt.Println("Write in temp file")
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	// write this byte array to our temporary file
	tempFileWriter := bufio.NewWriter(tempFile)
	tempFileWriter.Write(fileBytes)
	tempFileWriter.Flush()

	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")

	//defer tempFile.Close()
	fileToParser, _ := os.Open(tempFile.Name())
	processFile(fileToParser)

}

func consumeChainParcerCsv(db *sql.DB) {

	for {
		company, more := <-chanParceCsv
		if more {
			insertCompany(db, company)
		} else {
			db.Close()
			return
		}
	}

}

func insertCompany(db *sql.DB, company Companies) {

	fmt.Println("inserting company")

	toJson, _ := json.Marshal(company)
	fmt.Println(string(toJson))

	company.Name = strings.ToUpper(company.Name)

	if len(company.Zip) != 5 {
		log.Println("invalid zipcode size, expected 5 but actual is ", len(company.Zip))
		return
	}

	_, err := db.Exec("insert into company (name, zipCode) values ($1, $2)", company.Name, company.Zip)
	if err != nil {
		panic(err)
	}

	fmt.Println("inserted company with")
}

func parseCsvFile(csvFile *os.File) {
	fmt.Println("parseCsvFile ", csvFile.Name())

	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ';'
	firstLine := true
	for {
		line, error := reader.Read()
		fmt.Println("foi aqui", line)
		if error == io.EOF {
			fmt.Println("EOF in ", line)
			break
		} else if error != nil {
			fmt.Println("ERROR in ", line)
			log.Fatal(error)
			return
		}
		if firstLine {
			firstLine = false
		} else {
			var company = Companies{
				Name: line[0],
				Zip:  line[1],
			}
			fmt.Println("send to chanparceCsv")

			chanParceCsv <- company

		}

	}
	fmt.Println("finish parser file")

	chanDone <- true

}

type Companies struct {
	Name string
	Zip  string
}
