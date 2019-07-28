package main

import(
    "database/sql"
    "encoding/csv"
    "io"
    "github.com/joho/godotenv"
    "log"
    "strings"
    "os"
    "fmt"
    _ "github.com/lib/pq"
)

func main() {
        
    e := godotenv.Load()
    if e != nil {
        fmt.Print(e)
    }

    username := os.Getenv("db_user")
    password := os.Getenv("db_pass")
    dbName := os.Getenv("db_name")
    dbHost := os.Getenv("db_host")
    dbPort := os.Getenv("db_port")

    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s "+
        "sslmode=disable", dbHost, dbPort, username, password, dbName)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    delete_query := `DROP TABLE IF EXISTS companies;`

    _, err = db.Exec(delete_query)
    if err != nil {
        panic(err)
    }

    create_query := `
    CREATE TABLE IF NOT EXISTS companies (
            Id serial primary key,
            name varchar,
            zip varchar
    );`

    _, err = db.Exec(create_query)
    if err != nil {
        panic(err)
    }

    insert_query := `
    INSERT INTO companies (Id, name, zip)
    VALUES ($1, $2, $3);`

    csvFile, _ := os.Open("./q1_catalog.csv")
    reader := csv.NewReader(csvFile)
    reader.Comma = ';'
    id := -1 
    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        } else if id == -1 {
            id += 1
            continue
        }

        id += 1

        name := strings.ToUpper(line[0])
        zip := fmt.Sprintf("%05s", line[1])

        _, err = db.Exec(insert_query, id, name, zip)
        if err != nil {
            panic(err)
        }   
    }
}
